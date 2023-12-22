package updater

import (
	"fmt"
	"github.com/metacubex/mihomo/component/geodata"
	"github.com/metacubex/mihomo/component/mmdb"
	C "github.com/metacubex/mihomo/constant"
	"github.com/oschwald/maxminddb-golang"
)

func UpdateMMDB(path string) (err error) {
	defer mmdb.ReloadIP()
	data, err := downloadForBytes(C.MmdbUrl)
	if err != nil {
		return fmt.Errorf("can't download MMDB database file: %w", err)
	}
	instance, err := maxminddb.FromBytes(data)
	if err != nil {
		return fmt.Errorf("invalid MMDB database file: %s", err)
	}
	_ = instance.Close()

	mmdb.IPInstance().Reader.Close()
	if err = saveFile(data, path); err != nil {
		return fmt.Errorf("can't save MMDB database file: %w", err)
	}
	return nil
}

func UpdateASN(path string) (err error) {
	defer mmdb.ReloadASN()
	data, err := downloadForBytes(C.ASNUrl)
	if err != nil {
		return fmt.Errorf("can't download ASN database file: %w", err)
	}

	instance, err := maxminddb.FromBytes(data)
	if err != nil {
		return fmt.Errorf("invalid ASN database file: %s", err)
	}
	_ = instance.Close()

	mmdb.ASNInstance().Reader.Close()
	if err = saveFile(data, path); err != nil {
		return fmt.Errorf("can't save ASN database file: %w", err)
	}
	return nil
}

func UpdateGeoIp(path string) (err error) {
	geoLoader, err := geodata.GetGeoDataLoader("standard")
	data, err := downloadForBytes(C.GeoIpUrl)
	if err != nil {
		return fmt.Errorf("can't download GeoIP database file: %w", err)
	}
	if _, err = geoLoader.LoadIPByBytes(data, "cn"); err != nil {
		return fmt.Errorf("invalid GeoIP database file: %s", err)
	}
	if err = saveFile(data, path); err != nil {
		return fmt.Errorf("can't save GeoIP database file: %w", err)
	}
	return nil
}

func UpdateGeoSite(path string) (err error) {
	geoLoader, err := geodata.GetGeoDataLoader("standard")
	data, err := downloadForBytes(C.GeoSiteUrl)
	if err != nil {
		return fmt.Errorf("can't download GeoSite database file: %w", err)
	}

	if _, err = geoLoader.LoadSiteByBytes(data, "cn"); err != nil {
		return fmt.Errorf("invalid GeoSite database file: %s", err)
	}

	if err = saveFile(data, path); err != nil {
		return fmt.Errorf("can't save GeoSite database file: %w", err)
	}
	return nil
}
