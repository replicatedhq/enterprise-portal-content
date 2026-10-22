---
id: 2026-10-22-option-selector-and-tabs
title: OptionSelector and Tabs
published_at: 2026-10-22T18:00:00Z
impact: optional
summary: "OptionSelector remembers a choice for the browser session. Tabs stay local to the page. The default home and support bundle pages already follow that split."
affects:
  - home
  - support bundles
---

`<OptionSelector>` and `<Tabs>` are both valid. They do different jobs.

Use `<OptionSelector>` when the customer should keep a named choice after a reload in the same session. Pass `storageKey` so every selector on the page with that key restores the same saved value. Logout and switching teams clear it. The default home page uses this for Linux / Helm quick links:

```md
<OptionSelector label="Install Method" defaultOption="Linux" storageKey="install-method">
<Option value="Linux">
- [Linux Installation](/installation/linux)
</Option>
<Option value="Helm">
- [Helm Installation](/installation/helm)
</Option>
</OptionSelector>
```

Use `<Tabs>` when the sections are local to that page. Tabs do not persist and do not share a choice with `<OptionSelector>`. The default support bundle page wraps each method's collection component in entitlement-gated tabs. Keep that pattern. Do not put `storageKey="install-method"` on those tabs, and do not convert them to `<OptionSelector>` to share the home page's Linux / Helm picker. The titles and the set of methods are not the same choice.

```md
<Tabs>
{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
<Tab title="Linux (Embedded Cluster)">

<LinuxBundles />

</Tab>
{{/if}}
</Tabs>
```

The persist fix for `<OptionSelector>` ships in the portal. Existing content repos do not need to copy pages for this update. A repo created from the template already has the default home and support bundle pages above.

## Apply this update to your repo

No page checkout is required unless you want to copy the default home or support bundle pages from the template. If you already forked those files, keep `<OptionSelector>` on home and `<Tabs>` on support bundles unless your own content has a different split.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

### 3. Review, preview, commit, and push

Only if you are taking template pages:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/home.md pages/support/bundles.md
git commit -m "Adopt Enterprise Portal template update: OptionSelector and Tabs"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
