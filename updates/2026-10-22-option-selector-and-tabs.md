---
id: 2026-10-22-option-selector-and-tabs
title: Tabs groupId and OptionSelector
published_at: 2026-10-28T18:00:00Z
impact: optional
summary: "The default home page uses Tabs with groupId for the Linux and Helm quick links. OptionSelector stays supported. Existing repos do not need to change."
affects:
  - home
---

`<Tabs>` is the component to use when a customer should keep a tab for the browser session. Pass the same `groupId` on every `<Tabs>` that should share that choice. A click selects that title on every other `<Tabs>` with that `groupId` on the page, and the next page restores it. The match is the tab title. Logging out from the user menu and switching teams clear it. The Log out form on the signup page does not. Omit `groupId` when the tabs are local to one page.

The default home page uses this for the Linux / Helm quick links. This is the block in `pages/home.md`:

```md
<Tabs groupId="install-method">
<Tab title="Linux">

{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
- [Linux Installation](installation/linux)
{{/if}}
- [Release History](installation/release-history)
- [Instances & Updates](updates/instances)
- [Support Bundles](support/bundles)
- [FAQ](support/faq)

</Tab>
<Tab title="Helm">

{{#if entitlements.isHelmInstallEnabled}}
- [Helm Installation](installation/helm)
{{/if}}
- [Release History](installation/release-history)
- [Instances & Updates](updates/instances)
- [Support Bundles](support/bundles)
- [FAQ](support/faq)

</Tab>
</Tabs>
```

`<OptionSelector>` stays supported. It keeps its current look. `storageKey` is the same session key as `groupId`, so an existing selector keeps working, including two selectors on one page that share a `storageKey`. New home pages should use `<Tabs>` with `groupId`. There is no removal date for `<OptionSelector>`.

The default support bundle page stays on `<Tabs>` with no `groupId`. Those titles are entitlement-gated, and they are not the home page's Linux / Helm choice. Do not add `groupId="install-method"` to those tabs.

A repo created from the template after this update already has the home page above. An existing repo keeps `<OptionSelector>` until you change that page. No page checkout is required.

`groupId` is ignored on a portal build that does not include it yet. Wait for that build before you copy the home page markup into a repo customers are using.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Skip this if the existing home page should stay as it is. To match the template, add the template as `upstream` if it is not already there, fetch, then take `pages/home.md` from this update's template commit or copy the Tabs block into your home page.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `15315d04d9a5e7ba2cc218c3bf5cc1307d8940e9`.

### 3. Compare your home page

```shell
git diff HEAD 15315d04d9a5e7ba2cc218c3bf5cc1307d8940e9 -- pages/home.md
```

### 4. Take the new page, or copy the Tabs block

If you have not customized `pages/home.md`:

```shell
git checkout 15315d04d9a5e7ba2cc218c3bf5cc1307d8940e9 -- pages/home.md
```

If you keep a customized home page, replace the quick-link selector with the Tabs block above. The install links are relative, and they sit inside the entitlement checks.

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/home.md
git commit -m "Adopt Enterprise Portal template update: Tabs groupId on home"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
