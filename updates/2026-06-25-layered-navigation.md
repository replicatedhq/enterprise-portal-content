---
id: 2026-06-25-layered-navigation
title: Layered navigation support
published_at: 2026-06-25T18:59:25Z
impact: optional
summary: Enterprise Portal now supports nested table-of-contents entries so vendors can group related pages under expandable sidebar sections.
affects:
  - navigation
  - table of contents
  - content organization
---

Enterprise Portal content can now use nested `items` in `toc.yaml`. This lets you organize related pages into layered sidebar groups instead of flattening every page into the top-level section.

Existing flat navigation continues to work. Adopt this update when your portal has enough pages that grouping improves scanability.

## What changed

- Sidebar items can now contain other sidebar items recursively.
- Parent items can link to a page and still expand to show child pages.
- Deep links automatically open their parent navigation groups.
- Helm reference anchor links remain grouped under their chart page.
- Sidebar open state and scroll position persist while customers browse.

## Example

Use nested `items` under any navigation item:

```yaml
navigation:
  - title: Installation
    icon: rocket
    items:
      - title: Requirements
        page: pages/installation/requirements.md
      - title: Linux
        page: pages/installation/linux.md
        items:
          - title: System Requirements
            page: pages/installation/linux/system-requirements.md
          - title: Advanced Configuration
            page: pages/installation/linux/advanced-configuration.md
      - title: Helm
        page: pages/installation/helm.md
```

## Recommended adoption

Review your `toc.yaml` and group pages that naturally belong together, such as advanced install options, configuration references, troubleshooting pages, or product-specific guides. Keep high-traffic pages near the top of each section so the first level of the sidebar remains easy to scan.

## Apply this update to your repo

Vendor content repos are **cloned** from the Enterprise Portal template — they are not forks — so a plain `git pull` or `git merge` from the template will not work. The correct approach is to add the template as an upstream remote, fetch its changes, and then review what differs.

### 1. Set up the upstream remote (one-time)

Run this once per content repo. If `upstream` already exists, skip to step 2.

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

### 3. Compare your table of contents

This update introduces nested `items` in `toc.yaml`. The template does not add new pages, but reviewing the upstream `toc.yaml` helps you spot natural groupings:

```shell
git diff HEAD upstream/main -- toc.yaml
```

### 4. Edit your table of contents

Open your local `toc.yaml` and introduce nested `items` where pages belong together. Use the YAML example from this guide as a reference for the syntax. This is a manual edit — do not blindly check out `toc.yaml` from upstream, as it will overwrite your custom navigation.

### 5. Review, preview, commit, and push

Inspect the changes, run a local preview, then commit and push:

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add toc.yaml
git commit -m "Adopt Enterprise Portal template update: layered navigation"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. You can find your app's preview command in Enterprise Portal > Content > Preview.
