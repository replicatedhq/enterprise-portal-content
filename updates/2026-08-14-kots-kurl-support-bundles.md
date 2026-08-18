---
id: 2026-08-14-kots-kurl-support-bundles
title: Support bundle collection for KOTS and kURL
published_at: 2026-08-14T00:00:00Z
impact: recommended
summary: The support bundle page now covers all four install methods, with collection steps for KOTS and kURL supplied by built-in components and each tab gated on the customer's entitlements.
affects:
  - support bundles
  - entitlements
---

The support bundle page previously offered collection steps for Embedded Cluster and Helm only. Customers entitled to KOTS or kURL saw no steps for their install method.

Two new components supply those steps: `<KurlBundles />` and `<KotsBundles />`. They match the existing `<LinuxBundles />` and `<HelmBundles />`, so the page composes all four the same way. Each tab is wrapped in the entitlement for its install method, so a customer sees only the methods their license covers:

```markdown
<Tabs>
{{#if entitlements.isKurlInstallEnabled}}
<Tab title="Linux (kURL)">

<KurlBundles />

</Tab>
{{/if}}
</Tabs>
```

The four conditions are `isEmbeddedClusterDownloadEnabled`, `isKurlInstallEnabled`, `isHelmInstallEnabled`, and `isKotsInstallEnabled`. None of them require airgap support, because these are collection instructions rather than airgap downloads.

The Embedded Cluster and Helm tabs are retitled from `Linux` and `Helm` to `Linux (Embedded Cluster)` and `Existing Cluster (Helm)`, matching the install-track names already used in the Installation section. Tab titles are yours to change; the steps inside them come from the components. Those steps are also laid out differently now, with a short lead paragraph and a labelled heading per step or path in place of the `UI based collection` and `CLI based collection` headings.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. A repo created from a template starts with its own commit history rather than a copy of the template's, so the two share no common ancestor and a plain `git pull` or `git merge` from the template fails with `refusing to merge unrelated histories`. The correct approach is to add the template as an upstream remote, fetch its changes, and then review what differs.

### 1. Set up the upstream remote (one-time)

Run this once per content repo. If `upstream` already exists, skip to step 2.

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `491625dc76b9c661b7a73cb638018c19ddcacfbc`, so later template changes are not pulled in accidentally.

Before checking out template files, make sure you do not have uncommitted work in the target paths:

```shell
git status --short
```

Commit or stash any local changes before continuing. The restore command below returns files to `HEAD`; it cannot recover uncommitted edits overwritten by checkout.

### 3. Compare your support bundle page

This update touches one file, `pages/support/bundles.md`. No pages were added or removed, and `toc.yaml` is unchanged.

```shell
git diff HEAD 491625dc76b9c661b7a73cb638018c19ddcacfbc -- pages/support/bundles.md
```

### 4. Take the new page

If the diff in step 3 shows nothing you want to keep, take the template's version of the page:

```shell
git checkout 491625dc76b9c661b7a73cb638018c19ddcacfbc -- pages/support/bundles.md
```

If you started from a clean worktree and decide not to keep the copied file, restore it before committing:

```shell
git checkout HEAD -- pages/support/bundles.md
```

This is the recommended route. Everything on this page other than the tab titles, the tip at the top, and the section headings is a component tag, so for most repos the template's version is simply the page you want.

Note that this replaces the whole file, not just the tabs. If the diff shows customizations you want to preserve, make the three changes by hand instead:

- **The two new components.** Add `<KurlBundles />` and `<KotsBundles />`, each in its own `<Tab>`, alongside the `<LinuxBundles />` and `<HelmBundles />` tabs you already have.
- **The entitlement conditions.** Wrap each tab in the condition for its install method, so a customer sees only the methods their license covers. The template orders the two Linux tabs before the two existing-cluster tabs.
- **The tab titles, if you want them.** The template retitles `Linux` to `Linux (Embedded Cluster)` and `Helm` to `Existing Cluster (Helm)`. If you prefer your own titles, keep them; the steps inside each tab come from the components either way.

### 5. Review, preview, commit, and push

Inspect the changes, run a local preview, then commit and push. Use `git diff HEAD` rather than a bare `git diff`, because the checkout in step 4 stages the change and a bare `git diff` would show nothing:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/support/bundles.md
git commit -m "Adopt Enterprise Portal template update: KOTS and kURL support bundle collection"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. You can find your app's preview command in Enterprise Portal > Content > Preview.
>
> If you maintain version branches, apply and push this update on each branch where customers should see it.
