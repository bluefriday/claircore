package rhel

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/quay/claircore"
	"github.com/quay/claircore/pkg/ovalutil"
	"github.com/quay/goval-parser/oval"
	"github.com/rs/zerolog"
)

func (u *Updater) Parse(ctx context.Context, r io.ReadCloser) ([]*claircore.Vulnerability, error) {
	log := zerolog.Ctx(ctx).With().
		Str("component", "rhel/Updater.Parse").
		Logger()
	ctx = log.WithContext(ctx)
	log.Info().Msg("starting parse")
	defer r.Close()
	root := oval.Root{}
	if err := xml.NewDecoder(r).Decode(&root); err != nil {
		return nil, fmt.Errorf("rhel: unable to decode OVAL document: %w", err)
	}
	log.Debug().Msg("xml decoded")
	protoVuln := func(def oval.Definition) (*claircore.Vulnerability, error) {
		return &claircore.Vulnerability{
			Updater:            u.Name(),
			Name:               def.Title,
			Description:        def.Description,
			Links:              ovalutil.Links(def),
			Severity:           def.Advisory.Severity,
			NormalizedSeverity: NormalizeSeverity(def.Advisory.Severity),
			// each updater is configured to parse a rhel release
			// specific xml database. we'll use the updater's release
			// to map the parsed vulnerabilities
			Dist: releaseToDist(u.release),
		}, nil
	}
	vulns, err := ovalutil.RPMDefsToVulns(ctx, root, protoVuln)
	if err != nil {
		return nil, err
	}
	return vulns, nil
}
