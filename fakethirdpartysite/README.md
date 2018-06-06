# fake site

tl;dr; the docker image and everything else in this subdirectory adds nothing to the purpose of the demo. Feel free to ignore ./*

Originally this demo simply scraped from the original source - the idea of this being that real network disruption in the container would have a visible effect on monitoring. Quickly the "conference-wifi" fear grew on me and I think using a real internet source for a live demo is likely to end in misery. Thus, this subdirectory exists to serve some of the original content in a local container that will be linked to the go apps. Network disruption in the goapps should have the same affect as previously desired.
