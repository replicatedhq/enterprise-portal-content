---
id: 2026-09-28-kots-kurl-upgrade-guides
title: KOTS and kURL application upgrade guides
published_at: 2026-09-28T23:25:00Z
impact: recommended
summary: The default Instances & Updates section now links KOTS and kURL instance cards to application upgrade guides. Copy the new pages and navigation or customers will not see View full upgrade guide. Adopt this only after the portal build that accepts kotsGuideHref and kurlGuideHref (vandoor #10558) is deployed.
affects:
  - instances
  - kots
  - kurl
---

The default template adds two guides and points the instance cards at them.

- `pages/updates/kots.md` is the Admin Console procedure for an application installed with KOTS. The reopen command uses `--namespace {{app.slug}}`.
- `pages/updates/kurl.md` is the application procedure for kURL. The reopen command uses `--namespace default`. It tells the customer not to rerun `install.sh` just because an application release is available. Cluster maintenance stays a procedure you document.
- `pages/updates/instances.md` passes `kotsGuideHref="/updates/kots"` and `kurlGuideHref="/updates/kurl"` to `<InstancesAndUpdates />`.
- `toc.yaml` lists Instances, Upgrade KOTS Applications, and Upgrade kURL Applications under Instances & Updates. The KOTS item is gated by `isKotsInstallEnabled`. The kURL item is gated by `isKurlInstallEnabled`.

A content repo created from the template after this update already has the pages. An existing repo keeps its old Instances page until you copy them.

Do not adopt this until vandoor #10558 is deployed. An older portal drops `kotsGuideHref` and `kurlGuideHref`, and the guide links do not render.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Add the template as `upstream` if it is not already there, fetch, then take the pages from this update's template commit. If you have customized `toc.yaml`, add the navigation entries yourself instead of replacing the file.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `cd91b815e1eb0d62d96b18248bdd4f71141b8e9c`.

### 3. Compare your pages

```shell
git diff HEAD cd91b815e1eb0d62d96b18248bdd4f71141b8e9c -- pages/updates/instances.md pages/updates/kots.md pages/updates/kurl.md toc.yaml
```

### 4. Take the new pages

If you have not customized `pages/updates/instances.md`:

```shell
git checkout cd91b815e1eb0d62d96b18248bdd4f71141b8e9c -- pages/updates/instances.md pages/updates/kots.md pages/updates/kurl.md
```

If you keep a customized Instances page, add these props to `<InstancesAndUpdates />` and keep your other copy:

```md
<InstancesAndUpdates
  kotsGuideHref="/updates/kots"
  kurlGuideHref="/updates/kurl"
/>
```

Then check out the two new guides:

```shell
git checkout cd91b815e1eb0d62d96b18248bdd4f71141b8e9c -- pages/updates/kots.md pages/updates/kurl.md
```

Under Instances & Updates in `toc.yaml`, use child items instead of a single page entry:

```yaml
  - title: Instances & Updates
    items:
      - title: Instances
        page: pages/updates/instances.md
      - title: Upgrade KOTS Applications
        page: pages/updates/kots.md
        visible_when:
          entitlements:
            - isKotsInstallEnabled
      - title: Upgrade kURL Applications
        page: pages/updates/kurl.md
        visible_when:
          entitlements:
            - isKurlInstallEnabled
```

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/updates/instances.md pages/updates/kots.md pages/updates/kurl.md toc.yaml
git commit -m "Adopt Enterprise Portal template update: KOTS and kURL upgrade guides"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
