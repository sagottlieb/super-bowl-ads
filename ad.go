package main

type ad struct {
	Brand   string
	Title   string
	AvgRank float32
	AirTime string
	Link    string
}

/*
Sample of relevant html
<article id="post-" class="commercial-block collapsible-block filterable__item commercial-block--results" data-id="336283" data-shorturl="" data-quarter="2" data-advertiser="rocket-mortgage">
	<div class="collapsible-block__header commercial-block__header">
		<h3 class="commercial-block__title">
			<span class="commercial-block__rank">1.</span>
			<span class="commercial-block__category">Rocket Mortgage			</span>
			<a href="https://admeter.usatoday.com/commercials/certain-is-better-tracy-morgan-dave-bautista-liza-koshy/" class="commercial-block__video-title" title="Certain Is Better – Tracy Morgan, Dave Bautista &amp; Liza Koshy" itemprop="url">
				Certain Is Better – Tracy Morgan, Dave Bautista &amp; Liza Koshy			</a>
		</h3>
		<dl class="commercial-block__ranking-meta">
			<dt class="average-score">
				Avg. Score				</dt>
			<dd class="average-score__num">
				7.4				</dd>
		</dl>
	</div>
</article>
*/

var query = `
LET doc = DOCUMENT("https://admeter.usatoday.com/results/2021")

FOR ad IN ELEMENTS(doc, '#post-')
    LET link = ELEMENT(ad, 'a')
    RETURN {
		link: link.attributes.href,
		title: link.attributes.title,
	}
`
