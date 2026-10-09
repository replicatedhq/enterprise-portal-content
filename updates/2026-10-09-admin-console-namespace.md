---
id: 2026-10-09-admin-console-namespace
title: Vendor-configured Admin Console namespace
published_at: 2026-10-09T18:00:00Z
impact: recommended
summary: "KOTS and kURL Admin Console reopen commands accept a namespace MDX prop. Copy the generated command on the default upgrade guides so cards and the sample stay in sync. Adopt this only after vandoor #10637 is deployed."
affects:
  - instances
  - kots
  - kurl
---

The default KOTS and kURL upgrade guides mount `<AdminConsoleCommand />` for the reopen sample. With no `namespace` prop, kURL uses `default` and KOTS uses the app slug.

Set the same `namespace` on `<InstancesAndUpdates />` (instance cards) and on each `<AdminConsoleCommand />` in the guides so customers get that kotsadm namespace without rewriting the kubectl line.

```md
<InstancesAndUpdates
  kotsGuideHref="/updates/kots"
  kurlGuideHref="/updates/kurl"
  namespace="kotsadm"
/>
```

```md
<AdminConsoleCommand installType="kots" namespace="kotsadm" lead="Reopen the KOTS Admin Console by running" />
```

Do not adopt this until vandoor #10637 is deployed. An older portal drops the `namespace` prop and renders `<AdminConsoleCommand />` as inert HTML.

A content repo created from the template after this update already has the generated command. An existing repo keeps the hardcoded kubectl sample until you copy the guide pages.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Add the template as `upstream` if it is not already there, fetch, then take the pages from this update's template commit. If you have customized the guides, add `namespace` to your existing mounts instead of replacing the files.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `a3c22eaadc1ca5ee6736a476ab4b28b4b953364d`.

### 3. Compare your pages

```shell
git diff HEAD a3c22eaadc1ca5ee6736a476ab4b28b4b953364d -- pages/updates/kots.md pages/updates/kurl.md
```

### 4. Take the new pages

If you have not customized the upgrade guides:

```shell
git checkout a3c22eaadc1ca5ee6736a476ab4b28b4b953364d -- pages/updates/kots.md pages/updates/kurl.md
```

If you keep customized guides, replace the hardcoded `kubectl kots admin-console --namespace` lines with `<AdminConsoleCommand />` as in the template. Add the same `namespace` on `<InstancesAndUpdates />` in `pages/updates/instances.md` and on each `<AdminConsoleCommand />` when the Admin Console is not in `default` (kURL) or the app slug (KOTS).

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/updates/kots.md pages/updates/kurl.md
git commit -m "Adopt Enterprise Portal template update: Admin Console namespace"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
