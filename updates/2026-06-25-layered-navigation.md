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
