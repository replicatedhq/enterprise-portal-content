---
id: 2026-09-21-kots-online-cli-downloads
title: Optional workstation CLIs on the default KOTS online page
published_at: 2026-09-21T00:00:00Z
impact: recommended
summary: The default Existing Cluster (KOTS) online and proxy path now includes a CLI-only download list (KOTS, Preflight, Support Bundle) after the generated install commands. Air gap downloads stay in the air gap branch. Do not unwrap the air gap tag.
affects:
  - installation
  - kots
---

Online KOTS install still uses generated `curl https://kots.io/install` plus `kubectl kots install`. That does not offer Preflight or Support Bundle CLIs. The default page now mounts `<KotsDownloadAssets cliOnly={true} />` on the online/proxy branch so those workstation CLIs (and a version-pinned KOTS CLI) are available without showing the Admin Console bundle or the application air gap bundle.

The air gap branch is unchanged: full `<KotsDownloadAssets />` inside `<WhenNetwork mode="airgap">`. Do not move that tag out of the wrap. `cliOnly` keys off the mount, so a dual-entitled customer who selected **online** does not get air-gap installer files in that flow.

Your portal must include the `cliOnly` prop (vandoor follow-up to #10547). If you adopt this page before that release, the extra prop is ignored or the tag fails depending on your build; wait for the portal release that lists `cliOnly` on `KotsDownloadAssets`.
