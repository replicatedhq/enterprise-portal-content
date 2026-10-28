---
id: 2026-10-22-option-selector-and-tabs
title: Tabs groupId and OptionSelector
published_at: 2026-10-28T18:00:00Z
impact: optional
summary: "The default home page uses Tabs with groupId for the Linux and Helm quick links. OptionSelector stays supported. Existing repos do not need to change."
affects:
  - home
  - support bundles
---

`<Tabs>` is the component to use when a customer should keep a tab after a reload in the same session. Pass the same `groupId` on every `<Tabs>` that should restore that choice. Logout and switching teams clear it. Omit `groupId` when the tabs are local to one page.

The default home page uses this for the Linux / Helm quick links:

```md
<Tabs groupId="install-method">
<Tab title="Linux">
- [Linux Installation](/installation/linux)
</Tab>
<Tab title="Helm">
- [Helm Installation](/installation/helm)
</Tab>
</Tabs>
```

`<OptionSelector>` stays supported. It keeps its current look. `storageKey` is the same session key as `groupId`, so an existing selector keeps working, including two selectors on one page that share a `storageKey`. New home pages should use `<Tabs>` with `groupId`. There is no removal date for `<OptionSelector>`.

The default support bundle page stays on `<Tabs>` with no `groupId`. Those titles are entitlement-gated, and they are not the home page's Linux / Helm choice. Do not add `groupId="install-method"` to those tabs.

A repo created from the template after this update already has the home page above. An existing repo keeps `<OptionSelector>` until you change that page. No page checkout is required.

`groupId` is ignored on a portal build that does not include it yet. Wait for that build before you copy the home page markup into a repo customers are using.

## Apply this update to your repo

No page checkout is required. Copy the home page only if you want an existing repo to match the new template.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

### 3. Review, preview, commit, and push

Only if you are replacing the home page selector:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/home.md
git commit -m "Adopt Enterprise Portal template update: Tabs groupId on home"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
