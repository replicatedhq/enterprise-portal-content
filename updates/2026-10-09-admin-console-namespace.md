---
id: 2026-10-09-admin-console-namespace
title: Vendor-configured Admin Console namespace
published_at: 2026-10-09T18:00:00Z
impact: recommended
summary: "KOTS and kURL Admin Console reopen commands accept separate namespace MDX props. Copy the generated command on the default upgrade guides so cards and the sample stay in sync."
affects:
  - instances
  - kots
  - kurl
---

The default KOTS and kURL upgrade guides mount `<AdminConsoleCommand />` for the reopen sample. With no `namespace` prop, kURL uses `default` and KOTS uses the app slug.

Set `kotsNamespace` on `<InstancesAndUpdates />` and `namespace` on the matching `<AdminConsoleCommand />` only if every customer's Admin Console runs in that namespace, for example because your own install instructions tell customers to use it. Leave `kurlNamespace` unset unless you forked `install.sh` and every kURL customer's Admin Console runs in that namespace. These props do not change the namespace in the portal's KOTS install command. After you set one, the reopen command no longer tells the customer to substitute a different namespace.

```md
<InstancesAndUpdates
  kotsGuideHref="/updates/kots"
  kurlGuideHref="/updates/kurl"
  kotsNamespace="kotsadm"
/>
```

```md
<AdminConsoleCommand installType="kots" namespace="kotsadm" lead="Reopen the KOTS Admin Console by running" />
```

A content repo created from the template after this update already has the generated command. An existing repo keeps the hardcoded kubectl sample until you copy the guide pages. If you mount `<AdminConsoleCommand />` on a portal that does not include that component, `/updates/kots` and `/updates/kurl` fail to render.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Add the template as `upstream` if it is not already there, fetch, then take the pages from this update's template commit. If you have customized the guides, add `kotsNamespace` / `kurlNamespace` and each `<AdminConsoleCommand />` `namespace` to your existing mounts instead of replacing the files.

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

If you keep customized guides, replace the hardcoded `kubectl kots admin-console --namespace` lines with `<AdminConsoleCommand />` as in the template. Add `kotsNamespace` or `kurlNamespace` on `<InstancesAndUpdates />` in `pages/updates/instances.md`, and `namespace` on each `<AdminConsoleCommand />`, when that type's Admin Console is not in `default` (kURL) or the app slug (KOTS).

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/updates/kots.md pages/updates/kurl.md pages/updates/instances.md
git commit -m "Adopt Enterprise Portal template update: Admin Console namespace"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
