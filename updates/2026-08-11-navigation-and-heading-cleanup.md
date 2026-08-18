---
id: 2026-08-11-navigation-and-heading-cleanup
title: Support bundle headings and sidebar icons
published_at: 2026-08-11T00:00:00Z
impact: recommended
summary: Two component changes affect every portal: support bundle sections no longer carry their own headings, and sidebar section icons are now opt-in rather than defaulted. Both may need a small change to your content. The default template also drops the Requirements page and reorders some sections.
affects:
  - support bundles
  - sidebar
  - navigation
  - installation
---

Two changes below already affect your portal, whether or not you touch your content repo: support bundle headings and sidebar icons. The third is a correction worth making to your own copy. Everything after that is a change to the default template, relevant only if you follow it.

## Add a heading above your support bundle collection tabs

`<LinuxBundles />` and `<HelmBundles />` used to render a `Support Bundle Collection` heading of their own. They no longer do, because a page that already titled the section ended up with two headings for it.

The template shipped without a heading there, relying on the component to supply one. If your content repo still looks like that, the collection section is now untitled while the sections below it keep their headings. Add one above the tabs:

```mdx
## Generate a Bundle

<Tabs>
<Tab title="Linux">
```

`<SupportBundleUpload />` also stops rendering its own heading, but the template already labels that section with `## Upload an Existing Bundle`, so most repos need no change there. `<InstancesAndUpdates />` likewise stops rendering its own title and lede, which `pages/updates/instances.md` already supplies.

## Check your sidebar section icons

Sidebar section icons are now opt-in. A section in your `toc.yaml` that omits `icon:` previously fell back to a generic book glyph. It now renders as text only.

If icons disappeared from your sidebar, that is why. Sections that already specify `icon:` are unaffected. To keep a glyph, set it explicitly:

```yaml
- title: Installation
  icon: rocket
  items:
    - title: Release History
      page: pages/installation/release-history.md
```

The sidebar was also restyled in the same pass. Section headings sit at full contrast, and each section's pages are grouped against a guide rail that marks the page you are on. That applies automatically and needs no change to your content.

## Correct an inaccurate support bundle claim

A note on the support bundle page stated that bundles "do not include secrets or sensitive data" and can be safely shared. Redaction is spec-driven and best-effort, so that is not a guarantee, and it contradicted the collection instructions that recommend reviewing the archive before uploading. The note has been removed.

If your content repo still carries that wording, removing it is worth doing regardless of whether you adopt anything else in this update.

## Changes to the default template

These affect the template only. Adopt them if you follow the default content; ignore them if you have your own structure.

- **The Requirements page is gone.** `pages/installation/requirements.md` has been removed and its prerequisites now sit directly on the Embedded Cluster installation page, where customers act on them. Inbound links from `pages/home.md`, `pages/support/faq.md`, and `pages/installation/linux.md` were removed with it. If you have customized that page, keep it.
- **Automation moved to the end of the sidebar.** It is a reference section, not a starting point.
- **Upload now comes before history on the support bundle page.** `Upload an Existing Bundle` precedes `Uploaded Bundles`, so uploading comes before the list of past uploads.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. A repo created from a template starts with its own commit history rather than a copy of the template's, so a plain `git pull` or `git merge` from the template will not work. Add the template as an upstream remote, fetch its changes, and adopt selectively.

### 1. Set up the upstream remote (one-time)

Run this once per content repo. If `upstream` already exists, skip to step 2.

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `f2b916cde77a760123a7895b67ac9b4145815fba`, so later template changes are not pulled in accidentally.

### 3. Adopt the content changes

This update touches several files. Adopt them selectively based on your repo's customizations.

**Lower-risk to copy from upstream** — if these files are still close to the template, you can copy them directly and then review the diff:

```shell
git checkout f2b916cde77a760123a7895b67ac9b4145815fba -- pages/support/bundles.md pages/updates/instances.md
```

If the checkout overwrites a file you want to keep, restore it before committing:

```shell
git checkout HEAD -- pages/support/bundles.md pages/updates/instances.md
```

**Review before copying** — these files are likely customized in vendor repos:

- **`pages/support/faq.md`** — The template removed the Requirements accordion. If your content repo still links to the removed requirements page, remove that accordion or update it to point to the prerequisites on the Embedded Cluster installation page.
- **`pages/installation/linux.md`** — Inbound links to the removed Requirements page were deleted from this file. Check for links such as `requirements` or `installation/requirements`. If present, remove those links or update them to point to the prerequisites on the Embedded Cluster installation page.
- **`pages/home.md`** — Inbound links to the removed Requirements page were deleted from this file. Check for links such as `requirements` or `installation/requirements`. If present, remove those links or update them.
- **`toc.yaml`** — The Automation section was moved to the end of the sidebar, and sidebar icons are now opt-in. Compare your file to upstream and decide which changes to keep:
  ```shell
  git diff HEAD f2b916cde77a760123a7895b67ac9b4145815fba -- toc.yaml
  ```
  Merge the relevant changes manually rather than overwriting the file wholesale.

**A file was removed:** `pages/installation/requirements.md` has been deleted from the template. If your content repo still carries that file, keep it if you still need it; otherwise remove it.

### 4. Review, preview, commit, and push

Inspect the changes, run a local preview, then stage the files you adopted. Adjust the file list if you skipped or moved any of these template files:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/support/bundles.md pages/support/faq.md pages/updates/instances.md
git add pages/installation/linux.md pages/home.md toc.yaml
git status --short
```

If you remove the Requirements page, stage that deletion too before committing:

```shell
git rm pages/installation/requirements.md
```

Then commit and push:

```shell
git commit -m "Adopt Enterprise Portal template update: support bundle headings and sidebar icons"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. You can find your app's preview command in Enterprise Portal > Content > Preview.
>
> If you maintain version branches, apply and push this update on each branch where customers should see it.
