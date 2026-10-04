# GitHub Actions CI/CD reference

A fully commented reference workflow for GitHub Actions, written to be read
rather than copied blindly, plus a small Go service to exercise it against.

## What is here

| File | What it is |
|---|---|
| `.github/workflows/REFERENCE-full-cicd-template.yml` | The reference. Every block carries the reasoning for the setting above it. |
| `.github/workflows/ci.yml` | The trimmed version that actually runs on this repository. |
| `main.go`, `main_test.go` | A small Go service, so the pipeline has something real to build and test. |

## What the reference covers

- **Least-privilege `permissions`.** Every scope is listed and set to `none` by
  default, with a comment stating the one case that justifies raising it.
  `contents: read` is the floor, because `actions/checkout` cannot work below it.
- **`concurrency` groups**, so a second push cancels the first run instead of
  racing it to the same environment.
- **OIDC federation** for cloud authentication, so no long-lived cloud
  credential is ever stored as a secret. `id-token: write` is set only on the
  jobs that need it.
- **Job sequencing** — build, test, scan, then deploy, with the gates in the
  order where a failure is cheapest to fix.

## Using it

Copy `REFERENCE-full-cicd-template.yml` into your own `.github/workflows/`,
then delete what you do not need. It is deliberately over-commented: the
comments are the point, and they are meant to be removed once the decision
behind each one has been made consciously.

---

### Who maintains this

[RRR Solution Providers](https://www.rrrsolutionproviders.ca?utm_source=github&utm_medium=readme&utm_campaign=repos&utm_content=githubaction) — cloud, Kubernetes and platform
engineering, Toronto, Canada.

We publish our prices, which is unusual in consulting: [www.rrrsolutionproviders.ca/pricing](https://www.rrrsolutionproviders.ca/pricing?utm_source=github&utm_medium=readme&utm_campaign=repos&utm_content=githubaction).
If you want something like this built properly in your own estate, the smallest
way to start is a [five-day fixed-price audit](https://www.rrrsolutionproviders.ca/audit?utm_source=github&utm_medium=readme&utm_campaign=repos&utm_content=githubaction) — read-only access, and you
keep the written report whether or not you continue with us.

Issues and corrections are welcome.
