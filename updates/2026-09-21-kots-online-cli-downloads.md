---
id: 2026-09-21-kots-online-cli-downloads
title: Workstation CLIs on the default KOTS online page
published_at: 2026-09-21T00:00:00Z
impact: recommended
summary: The default Existing Cluster (KOTS) online and proxy path includes a CLI-only download list (KOTS, Preflight, Support Bundle) after the generated install commands. Air gap downloads stay in the air gap branch. Do not unwrap the air gap tag. Adopt this page only after the portal build that supports cliOnly (vandoor #10549).
affects:
  - installation
  - kots
---

Online KOTS install still uses generated `curl https://kots.io/install` plus `kubectl kots install`. That does not offer Preflight or Support Bundle CLIs. The default page now mounts `<KotsDownloadAssets cliOnly={true} />` on the online/proxy branch so those workstation CLIs (and a version-pinned KOTS CLI) are available without showing the Admin Console bundle or the application air gap bundle.

The air gap branch is unchanged: full `<KotsDownloadAssets />` inside `<WhenNetwork mode="airgap">`. Do not move that tag out of the wrap. `cliOnly` keys off the mount, so a dual-entitled customer who selected **online** does not get air-gap installer files in that flow.

Do not adopt this page until the portal lists `cliOnly` on `KotsDownloadAssets` (vandoor #10549). If you ship the page first, an older portal may strip the prop and show Admin Console and air-gap bundles on the online branch for customers whose license also has air gap.
