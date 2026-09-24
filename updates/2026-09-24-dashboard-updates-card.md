---
id: 2026-09-24-dashboard-updates-card
title: Updates card on the portal home page
published_at: 2026-09-24T16:50:00Z
impact: recommended
summary: The default home page now includes an updates card with online and air gap instance counts, and how many of those instances have an update available. Add it to an existing content repo or customers will not see the card.
affects:
  - home
  - instances
---

The default home page includes `<UpdatesCard />`. The card shows active online instances and air gap instances, and badges how many of each have an update available. "View updates" appears when a page in the portal uses `<InstancesAndUpdates />`. The template page for that is `pages/updates/instances.md`.

The card is part of the Enterprise Portal application. It renders only on a page that includes the tag. Do not pass props to it.

A content repo created from the template after this update already has the tag. An existing repo keeps its old home page until you add the tag there.

## Apply this update to your repo

Your content repo was created from the Enterprise Portal template repository, not forked from it. Add the template as `upstream` if it is not already there, fetch, then take `pages/home.md` from this update's template commit or add the tag yourself.

### 1. Set up the upstream remote (one-time)

```shell
git remote add upstream https://github.com/replicatedhq/enterprise-portal-content.git
```

### 2. Fetch the latest template changes

```shell
git fetch upstream
```

The commands below use this update's template commit, `540b62b2962d59497bd283caf1018c184b86e20d`.

### 3. Compare your home page

```shell
git diff HEAD 540b62b2962d59497bd283caf1018c184b86e20d -- pages/home.md
```

### 4. Take the new page, or add the tag

If you have not customized `pages/home.md`:

```shell
git checkout 540b62b2962d59497bd283caf1018c184b86e20d -- pages/home.md
```

If you keep a customized home page, add this line after the available-features block and before `## Getting Started`:

```md
<UpdatesCard />
```

### 5. Review, preview, commit, and push

```shell
git diff HEAD
replicated enterprise-portal preview . --app <your-app-slug>
git add pages/home.md
git commit -m "Adopt Enterprise Portal template update: home page updates card"
git push
```

> **Note:** Replace `<your-app-slug>` with your app's slug. If you maintain version branches, apply and push this update on each branch where customers should see it.
