---
name: sync-back-pr
description: Bring internal-only work from the private fork's `internal/main` into this public repo, replacing the internal SDK with the public one and stripping internal-only wiring. Use when asked to "sync back to upstream", "upstream our internal changes", "open a sync-back PR", "pull internal changes into the public repo", or similar.
---

# sync-back-pr — upstream `internal/main` into this repo

This repo (`nutanix-cloud-native/prism-go-client`) is the public one. A private fork,
`nutanix-cloud-native/internal-prism-go-client`, carries internal development on its
`internal/main` branch. Two flows move code between them:

| | direction | branch | base | lives in |
|---|---|---|---|---|
| `sync-pr` | public → internal | `issue/sync-main` | `internal/main` | the internal fork |
| `sync-back-pr` | internal → public | `chore/sync-from-internal` | `main` | here |

This is the higher-risk direction: it moves code out of a private repo into a public one, and it
applies a diff produced against internal's CI and dependency wiring. Everything below exists to
stop something leaking, breaking public CI, or weakening this repo's security posture.

## Before doing anything: check for an existing sync-back PR

```bash
gh pr list --repo nutanix-cloud-native/prism-go-client --state open \
  --json number,title,headRefName,baseRefName,url
```

If a PR from `chore/sync-from-internal` is already open, update that branch in place. Never open
a duplicate.

Then check what else is in flight. A sync-back is large and will conflict with any open PR
touching `go.mod` or `converged/`. If one exists, **stack on it** (`--base <that branch>`) rather
than racing it — merging either one first would leave the other with a `go.mod` conflict.

## Working safely

Use a scratch worktree so the diff can be shaped without disturbing a checkout:

```bash
git worktree add /tmp/sync-back main
cd /tmp/sync-back
git fetch <path-or-url-of-internal-fork> internal/main:internal-main
git checkout -b chore/sync-from-internal internal-main
```

Rewrite dependencies and strip internal-only files (below), then squash onto `main`.

## 1. Dependencies: public SDK only, on latest tags

Rewrite every internal SDK import to its public equivalent:

```
github.com/nutanix-core/ntnx-api-golang-sdk-internal/<mod>-go-client/v17
  ->  github.com/nutanix/ntnx-api-golang-clients/<mod>-go-client/v4
```

Then pin each module to its newest tag. Read tags from the source of truth, not the module proxy,
which can lag:

```bash
git ls-remote --tags https://github.com/nutanix/ntnx-api-golang-clients.git \
  | awk '{print $2}' | sed 's|refs/tags/||' | grep '^<mod>-go-client/' | sort -V | tail -5
```

**Gotcha that has bitten before:** if `go.mod` still holds `v17.x` versions under a `/v4` module
path, `go get` aborts with *"invalid: should be v4, not v17"* and applies **none** of the
requested pins — including ones for modules you did rewrite. It reads like a warning; it is a
total no-op, and the result is a sync silently carrying stale SDK versions. Fix the versions in
`go.mod` first, then `go get`, then verify every line:

```bash
grep ntnx-api go.mod    # confirm each module is on the tag you intended
```

Prefer stable tags. A prerelease is acceptable only where a needed package genuinely does not
exist in the stable tag — internal features reach public betas first. Verify rather than assume:

```bash
go mod tidy   # names the exact missing package, e.g.
              # multidomain-go-client/v4@v4.3.1 does not contain .../request/projects
```

If a beta is required, say so in the PR body and name the package that forces it.

## 2. No internal wiring may remain

```bash
grep -rn "ntnx-api-golang-sdk-internal\|nutanix-core" \
  --include="*.go" --include="*.mod" --include="*.sum" .
```

Must come back empty — `go.sum` included.

Also drop internal-only tooling that has no meaning here: the internal fork's `sync-pr` skill
(`.agents/skills/sync-pr` and its `.claude/skills` symlink) and `.claude/settings.local.json`.

**Do not delete this skill.** `sync-back-pr` lives in this repo. Depending on when the internal
fork last ran `sync-pr`, `internal/main` may not contain it, in which case the sync-back diff
will show it as a deletion. Keep it — same hazard class as the Codecov removal below.

## 3. `.github/` — the security-critical part

**Default: carry over no `.github/` changes at all.** Restore this repo's versions wholesale:

```bash
git checkout main -- .github/
```

Deviate only for a genuine CI fix that is not internal-specific, and call it out in the PR body.
Both directions of divergence are dangerous.

### Internal additions that must NOT land here

- `GOPRIVATE` / `GONOSUMDB` env vars — meaningless once dependencies are public.
- `actions/create-github-app-token` steps using `GHA_CHECKOUT_APP_ID` /
  `GHA_CHECKOUT_APP_PRIVATE_KEY`, and the
  `git config --global url."https://x-access-token:...@github.com/...".insteadOf` rewrites.
  These mint credentials for private orgs. This repo runs workflows for **fork PRs**;
  token-minting steps here are a credential-exposure risk.
- `if: github.repository == 'nutanix-cloud-native/internal-prism-go-client'` guards. These never
  match here, so the job silently never runs. A permanently-skipped required check is worse than
  a failing one — it looks green.
- `internal/main` and `internal/release-*` branch triggers.

### Internal removals that must NOT propagate here

A sync-back applies internal's *diff*, so a step deleted in the fork gets deleted here too.
Guard against this explicitly:

- **Codecov.** `codecov/codecov-action` with `CODECOV_TOKEN` is intentionally absent in the fork,
  because Codecov rejects tokenless uploads for private repos. It must stay here. Do not carry
  the deletion back.
- Any `EXPORT_RESULT` / coverage plumbing this repo relies on.
- This repo's `branches: [main, release-*]` triggers.

### Runner changes — never carry these back

CodeQL here runs on `ubuntu-latest` deliberately, with a comment saying so: untrusted fork code
must compile in an ephemeral GitHub-hosted sandbox, **not** on self-hosted infrastructure. The
internal fork moves it to `self-hosted-nutanix-docker-small`, which is fine for a private repo
where every contributor is trusted. Carrying that here would let a fork PR execute attacker code
on Nutanix-controlled runners. Combined with `pull_request_target`, that is a textbook
pwn-request.

**Rule:** jobs in this repo stay on GitHub-hosted runners. If a sync-back diff moves a job onto a
self-hosted runner, or introduces `pull_request_target` where this repo had `pull_request`, stop
— that is a security regression, not a sync.

Confirm nothing unexpected survived:

```bash
git diff main -- .github/    # expect empty
```

## 4. Commit shape

Squash into a single commit on top of `main`. The internal history carries internal ticket IDs
and internal-CI commits describing infrastructure that does not exist here, and a merge would
import commits whose content this skill then reverts — incoherent public history.

Squashing costs per-commit authorship, so preserve it explicitly:

```bash
git log --format='%an <%ae>' main..internal-main | sort -u \
  | grep -viE "noreply@github|dependabot|actions@github"
```

Add each as a `Co-authored-by:` trailer. This is a public repo; do not erase contributors.

Per global instructions, no Claude/Anthropic attribution in commit messages or PR descriptions.

## 5. Verify before opening the PR

```bash
go build ./... && go vet ./...
go test ./environment/... ./internal/... ./versionutils/... ./v4/... ./converged/
gofmt -l .    # this repo is not gofmt-clean; compare against main and only fix new offenders
```

`konnector` and `v3` tests need a local **keploy** server on `:6789` and fail without it, ending
in a nil-pointer panic rather than a clean assertion failure. That is environmental — check the
same tests on clean `main` before blaming the sync.

The `converged/v4` tests are live-cluster integration tests. Credentials come from
`~/.prism-dev.yaml` (or `$PRISM_DEV_CONFIG`), **not** from `NUTANIX_*` env vars —
`internal/testhelpers/creds.go` sets those itself from the YAML, so exporting them by hand does
nothing. If you cannot reach a cluster, say so plainly in the PR body rather than implying the
sync was verified end to end.

## 6. PR body

Title: `feat: sync internal fork changes using public API SDKs`

Include:

1. **What is being upstreamed** — the new public API surface, in a paragraph.
2. **SDK versions**, and any prerelease pin with the package that forces it.
3. **What was deliberately excluded** — internal CI, the `sync-pr` skill — and why.
4. **Anything a reviewer would otherwise miss** — call-convention changes, behavioural changes,
   inverted test assertions.
5. **Honest test status**, including what was not run.

## Checklist

- [ ] No duplicate sync-back PR; stacked on any conflicting in-flight PR
- [ ] Zero `nutanix-core` / `ntnx-api-golang-sdk-internal` references, `go.sum` included
- [ ] Every public SDK module on its newest appropriate tag, re-checked in `go.mod` after `go get`
- [ ] Prereleases justified by a genuinely missing package, and named in the PR body
- [ ] `git diff main -- .github/` is empty, or every deviation is justified
- [ ] Codecov and other public-only CI steps still present
- [ ] No job moved to a self-hosted runner; no new `pull_request_target`
- [ ] `sync-pr` and `.claude/settings.local.json` not carried over; this skill not deleted
- [ ] Single squashed commit with `Co-authored-by:` trailers for every human author
- [ ] Build, vet and offline tests pass; unrun tests disclosed in the PR body
