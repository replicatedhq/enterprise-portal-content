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

Two changes below already affect your portal, whether or not you touch your fork: support bundle headings and sidebar icons. The third is a correction worth making to your own copy. Everything after that is a change to the default template, relevant only if you follow it.

## Add a heading above your support bundle collection tabs

`<LinuxBundles />` and `<HelmBundles />` used to render a `Support Bundle Collection` heading of their own. They no longer do, because a page that already titled the section ended up with two headings for it.

The template shipped without a heading there, relying on the component to supply one. If your fork still looks like that, the collection section is now untitled while the sections below it keep their headings. Add one above the tabs:

```mdx
## Generate a Bundle

<Tabs>
<Tab title="Linux">
```

`<SupportBundleUpload />` also stops rendering its own heading, but the template already labels that section with `## Upload an Existing Bundle`, so most forks need no change there. `<InstancesAndUpdates />` likewise stops rendering its own title and lede, which `pages/updates/instances.md` already supplies.

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

If your fork still carries that wording, removing it is worth doing regardless of whether you adopt anything else in this update.

## Changes to the default template

These affect the template only. Adopt them if you follow the default content; ignore them if you have your own structure.

- **The Requirements page is gone.** `pages/installation/requirements.md` has been removed and its prerequisites now sit directly on the Embedded Cluster installation page, where customers act on them. Inbound links from `pages/home.md`, `pages/support/faq.md`, and `pages/installation/linux.md` were removed with it. If you have customized that page, keep it.
- **Automation moved to the end of the sidebar.** It is a reference section, not a starting point.
- **Upload now comes before history on the support bundle page.** `Upload an Existing Bundle` precedes `Uploaded Bundles`, so uploading comes before the list of past uploads.

## Apply this update to your repo

Vendor content repos are **cloned** from the Enterprise Portal template — they are not forks — so a plain `git pull` or `git merge` from the template will not work. Add the template as an upstream remote, fetch its changes, and adopt selectively.

### 1. Set up the upstream remote (one-time)

Run this once per content repo. If `upstream` already exists, skip to step 2.

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

### 3. Adopt the content changes

This update touches several files. Adopt them selectively based on your fork's customizations.

**Safe to copy from upstream** — these changes are straightforward and do not risk overwriting important customizations:

```shell
git checkout upstream/main -- pages/support/bundles.md pages/support/faq.md pages/updates/instances.md
```

**Review before copying** — these files are likely customized in vendor repos:

- **`pages/installation/linux.md`** — Inbound links to the removed `requirements.md` page were deleted from this file. Check whether your fork still references `requirements.md`. If so, remove those links or update them to point to the prerequisites on the Embedded Cluster installation page.
- **`pages/home.md`** — Inbound links to `requirements.md` were removed. If your fork's home page still links to it, update those links.
- **`toc.yaml`** — The Automation section was moved to the end of the sidebar, and sidebar icons are now opt-in. Compare your file to upstream and decide which changes to keep:
  ```shell
  git diff HEAD upstream/main -- toc.yaml
  ```
  Merge the relevant changes manually rather than overwriting the file wholesale.

**A file was removed:** `pages/installation/requirements.md` has been deleted from the template. If your fork still carries that file, review whether it is still needed and remove it if so.

### 4. Review, preview, commit, and push

Inspect the changes, run a local preview, then commit and push:

```shell
git diff
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/support/bundles.md pages/support/faq.md pages/updates/instances.md
git commit -m "Adopt Enterprise Portal template update: support bundle headings and sidebar icons"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. You can find your app's preview command in Enterprise Portal > Content > Preview.
