package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type City struct {
	Name       string
	Population int
}

func main() {
	f, err := os.Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := genRows(f)
	filterSmallCity := filterByMinPopulation(40000)

	// fan out pattern
	// more than one worker competing to consume the same channel

	//       __ worker1
	// rows /__ worker2
	//      \__ worker3
	//       __ workern...
	ur1 := upperCityName(filterSmallCity(rows))
	ur2 := upperCityName(filterSmallCity(rows))
	ur3 := upperCityName(filterSmallCity(rows))

	// fan in pattern consolidates the outputs from multiple channels into one
	//
	// worker1 ___
	// worker2 ___\ output
	// worker3 ___/

	for c := range fanIn(ur1, ur2, ur3) {
		log.Println(c)
	}
}

func upperCityName(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(out)
	}()
	return out
}

func filterByMinPopulation(min int) func(<-chan City) <-chan City {
	return func(cities <-chan City) <-chan City {
		out := make(chan City)
		go func() {
			for c := range cities {
				if c.Population > min {
					out <- City{Name: c.Name, Population: c.Population}
				}
			}
			close(out)
		}()
		return out
	}
}

func genRows(r io.Reader) chan City {
	out := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			populationInt, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: populationInt,
			}
		}
		close(out)
	}()
	return out
}

func fanIn(chans ...<-chan City) <-chan City {
	out := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, c := range chans {
		go func(city <-chan City) {
			for r := range city {
				out <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//
//-- cities.csv --
//"city","city_ascii","lat","lng","country","iso2","iso3","admin_name","capital","population","id"
//"Tokyo","Tokyo","35.6897","139.6922","Japan","JP","JPN","Tōkyō","primary","37977000","1392685764"
//"Jakarta","Jakarta","-6.2146","106.8451","Indonesia","ID","IDN","Jakarta","primary","34540000","1360771077"
//"Delhi","Delhi","28.6600","77.2300","India","IN","IND","Delhi","admin","29617000","1356872604"
//"Mumbai","Mumbai","18.9667","72.8333","India","IN","IND","Mahārāshtra","admin","23355000","1356226629"
//"Manila","Manila","14.5958","120.9772","Philippines","PH","PHL","Manila","primary","23088000","1608618140"
//"Shanghai","Shanghai","31.1667","121.4667","China","CN","CHN","Shanghai","admin","22120000","1156073548"
//"São Paulo","Sao Paulo","-23.5504","-46.6339","Brazil","BR","BRA","São Paulo","admin","22046000","1076532519"
//"Seoul","Seoul","37.5833","127.0000","Korea, South","KR","KOR","Seoul","primary","21794000","1410836482"
//"Mexico City","Mexico City","19.4333","-99.1333","Mexico","MX","MEX","Ciudad de México","primary","20996000","1484247881"
//"Guangzhou","Guangzhou","23.1288","113.2590","China","CN","CHN","Guangdong","admin","20902000","1156237133"
//"Beijing","Beijing","39.9050","116.3914","China","CN","CHN","Beijing","primary","19433000","1156228865"
//"Cairo","Cairo","30.0561","31.2394","Egypt","EG","EGY","Al Qāhirah","primary","19372000","1818253931"
//"New York","New York","40.6943","-73.9249","United States","US","USA","New York","","18713220","1840034016"
//"Kolkāta","Kolkata","22.5411","88.3378","India","IN","IND","West Bengal","admin","17560000","1356060520"
//"Moscow","Moscow","55.7558","37.6178","Russia","RU","RUS","Moskva","primary","17125000","1643318494"
//"Bangkok","Bangkok","13.7500","100.5167","Thailand","TH","THA","Krung Thep Maha Nakhon","primary","17066000","1764068610"
//"Buenos Aires","Buenos Aires","-34.5997","-58.3819","Argentina","AR","ARG","Buenos Aires, Ciudad Autónoma de","primary","16157000","1032717330"
//"Shenzhen","Shenzhen","22.5350","114.0540","China","CN","CHN","Guangdong","minor","15929000","1156158707"
//"Dhaka","Dhaka","23.7161","90.3961","Bangladesh","BD","BGD","Dhaka","primary","15443000","1050529279"
//"Lagos","Lagos","6.4500","3.4000","Nigeria","NG","NGA","Lagos","minor","15279000","1566593751"
//"Istanbul","Istanbul","41.0100","28.9603","Turkey","TR","TUR","İstanbul","admin","15154000","1792756324"
//"Ōsaka","Osaka","34.6936","135.5019","Japan","JP","JPN","Ōsaka","admin","14977000","1392419823"
//"Karachi","Karachi","24.8600","67.0100","Pakistan","PK","PAK","Sindh","admin","14835000","1586129469"
//"Bangalore","Bangalore","12.9699","77.5980","India","IN","IND","Karnātaka","admin","13707000","1356410365"
//"Tehran","Tehran","35.7000","51.4167","Iran","IR","IRN","Tehrān","primary","13633000","1364305026"
//"Kinshasa","Kinshasa","-4.3317","15.3139","Congo (Kinshasa)","CD","COD","Kinshasa","primary","13528000","1180000363"
//"Ho Chi Minh City","Ho Chi Minh City","10.8167","106.6333","Vietnam","VN","VNM","Hồ Chí Minh","admin","13312000","1704774326"
//"Los Angeles","Los Angeles","34.1139","-118.4068","United States","US","USA","California","","12750807","1840020491"
//"Rio de Janeiro","Rio de Janeiro","-22.9083","-43.1964","Brazil","BR","BRA","Rio de Janeiro","admin","12272000","1076887657"
//"Nanyang","Nanyang","32.9987","112.5292","China","CN","CHN","Henan","","12010000","1156192287"
//"Chennai","Chennai","13.0825","80.2750","India","IN","IND","Tamil Nādu","admin","11324000","1356374944"
//"Chengdu","Chengdu","30.6636","104.0667","China","CN","CHN","Sichuan","admin","11309000","1156421555"
//"Lahore","Lahore","31.5497","74.3436","Pakistan","PK","PAK","Punjab","admin","11021000","1586801463"
//"Paris","Paris","48.8566","2.3522","France","FR","FRA","Île-de-France","primary","11020000","1250015082"
//"London","London","51.5072","-0.1275","United Kingdom","GB","GBR","London, City of","primary","10979000","1826645935"
//"Linyi","Linyi","35.0606","118.3425","China","CN","CHN","Shandong","","10820000","1156086320"
//"Tianjin","Tianjin","39.1467","117.2056","China","CN","CHN","Tianjin","admin","10800000","1156174046"
//"Shijiazhuang","Shijiazhuang","38.0422","114.5086","China","CN","CHN","Hebei","admin","10784600","1156217541"
//"Baoding","Baoding","38.8671","115.4845","China","CN","CHN","Hebei","","10700000","1156256829"
//"Zhoukou","Zhoukou","33.6250","114.6418","China","CN","CHN","Henan","minor","9901000","1156272098"
//"Lima","Lima","-12.0500","-77.0333","Peru","PE","PER","Lima","primary","9848000","1604728603"
//"Hyderābād","Hyderabad","17.3667","78.4667","India","IN","IND","Telangana","admin","9746000","1356871768"
//"Bogotá","Bogota","4.6126","-74.0705","Colombia","CO","COL","Bogotá","primary","9464000","1170483426"
//"Weifang","Weifang","36.7167","119.1000","China","CN","CHN","Shandong","","9373000","1156912965"
//"Nagoya","Nagoya","35.1167","136.9333","Japan","JP","JPN","Aichi","admin","9113000","1392407472"
//"Wuhan","Wuhan","30.5872","114.2881","China","CN","CHN","Hubei","admin","8962000","1156117184"
//"Heze","Heze","35.2333","115.4333","China","CN","CHN","Shandong","minor","8750000","1156225008"
//"Ganzhou","Ganzhou","25.8292","114.9336","China","CN","CHN","Jiangxi","minor","8677600","1156832475"
//"Tongshan","Tongshan","34.2610","117.1859","China","CN","CHN","Jiangsu","minor","8669000","1156241678"
//"Chicago","Chicago","41.8373","-87.6862","United States","US","USA","Illinois","","8604203","1840000494"
//"Handan","Handan","36.6116","114.4894","China","CN","CHN","Hebei","","8499000","1156957080"
//"Luanda","Luanda","-8.8383","13.2344","Angola","AO","AGO","Luanda","primary","8417000","1024949724"
//"Fuyang","Fuyang","32.8986","115.8045","China","CN","CHN","Anhui","","8360000","1156248008"
//"Kuala Lumpur","Kuala Lumpur","3.1478","101.6953","Malaysia","MY","MYS","Kuala Lumpur","primary","8285000","1458988644"
//"Jining","Jining","35.4000","116.5667","China","CN","CHN","Shandong","","8023000","1156504601"
//"Dongguan","Dongguan","23.0475","113.7493","China","CN","CHN","Guangdong","minor","7981000","1156478242"
//"Hanoi","Hanoi","21.0245","105.8412","Vietnam","VN","VNM","Hà Nội","primary","7785000","1704413791"
//"Pune","Pune","18.5196","73.8553","India","IN","IND","Mahārāshtra","","7764000","1356081074"
//"Chongqing","Chongqing","29.5500","106.5069","China","CN","CHN","Chongqing","admin","7739000","1156936556"
//"Changchun","Changchun","43.9000","125.2000","China","CN","CHN","Jilin","admin","7674439","1156078103"
//"Zhumadian","Zhumadian","32.9773","114.0253","China","CN","CHN","Henan","","7640000","1156184822"
//"Ningbo","Ningbo","29.8750","121.5492","China","CN","CHN","Zhejiang","minor","7639000","1156170787"
//"Onitsha","Onitsha","6.1667","6.7833","Nigeria","NG","NGA","Anambra","minor","7635000","1566593483"
//"Nanjing","Nanjing","32.0500","118.7667","China","CN","CHN","Jiangsu","admin","7496000","1156644065"
//"Hefei","Hefei","31.8639","117.2808","China","CN","CHN","Anhui","admin","7457027","1156332710"
//"Ahmadābād","Ahmadabad","23.0300","72.5800","India","IN","IND","Gujarāt","minor","7410000","1356304381"
//"Hong Kong","Hong Kong","22.3050","114.1850","Hong Kong","HK","HKG","","","7347000","1344982653"
//"Khartoum","Khartoum","15.6031","32.5265","Sudan","SD","SDN","Khartoum","primary","7282000","1729268475"
//"Nantong","Nantong","31.9829","120.8873","China","CN","CHN","Jiangsu","minor","7282835","1156045365"
//"Yancheng","Yancheng","33.3936","120.1339","China","CN","CHN","Jiangsu","minor","7260240","1156995410"
//"Foshan","Foshan","23.0292","113.1056","China","CN","CHN","Guangdong","minor","7194311","1156738403"
//"Nanning","Nanning","22.8192","108.3150","China","CN","CHN","Guangxi","admin","7153300","1156605439"
//"Hengyang","Hengyang","26.8968","112.5857","China","CN","CHN","Hunan","","7148344","1156696884"
//"Xi’an","Xi'an","34.2667","108.9000","China","CN","CHN","Shaanxi","admin","7135000","1156244079"
//"Shenyang","Shenyang","41.8039","123.4258","China","CN","CHN","Liaoning","admin","7105000","1156309382"
//"Tangshan","Tangshan","39.6292","118.1742","China","CN","CHN","Hebei","","7100000","1156904299"
//"Shaoyang","Shaoyang","27.2418","111.4725","China","CN","CHN","Hunan","","7071000","1156310366"
//"Changsha","Changsha","28.1987","112.9709","China","CN","CHN","Hunan","admin","7044118","1156961497"
//"Santiago","Santiago","-33.4500","-70.6667","Chile","CL","CHL","Región Metropolitana","primary","7007000","1152554349"
//"Zhengzhou","Zhengzhou","34.7492","113.6605","China","CN","CHN","Henan","admin","7005000","1156183137"
//"Zhanjiang","Zhanjiang","21.1967","110.4031","China","CN","CHN","Guangdong","minor","6900000","1156502170"
//"Riyadh","Riyadh","24.6500","46.7100","Saudi Arabia","SA","SAU","Ar Riyāḑ","primary","6881000","1682999334"
//"Cangzhou","Cangzhou","38.3037","116.8452","China","CN","CHN","Hebei","","6800000","1156698069"
//"Dar es Salaam","Dar es Salaam","-6.8000","39.2833","Tanzania","TZ","TZA","Dar es Salaam","primary","6698000","1834843853"
//"Maoming","Maoming","21.6618","110.9178","China","CN","CHN","Guangdong","minor","6706000","1156568722"
//"Huanggang","Huanggang","30.4500","114.8750","China","CN","CHN","Hubei","minor","6667000","1156200037"
//"Xinyang","Xinyang","32.1264","114.0672","China","CN","CHN","Henan","","6634000","1156273453"
//"Shangrao","Shangrao","28.4419","117.9633","China","CN","CHN","Jiangxi","minor","6579714","1156405492"
//"Luoyang","Luoyang","34.6587","112.4245","China","CN","CHN","Henan","","6549941","1156069315"
//"Bijie","Bijie","27.3019","105.2863","China","CN","CHN","Guizhou","","6537498","1156018927"
//"Yantai","Yantai","37.3997","121.2664","China","CN","CHN","Shandong","","6516900","1156932948"
//"Quanzhou","Quanzhou","24.9139","118.5858","China","CN","CHN","Fujian","minor","6480000","1156829655"
//"Hangzhou","Hangzhou","30.2500","120.1675","China","CN","CHN","Zhejiang","admin","6446000","1156275243"
//"Miami","Miami","25.7839","-80.2102","United States","US","USA","Florida","","6445545","1840015149"
//"Kunming","Kunming","25.0433","102.7061","China","CN","CHN","Yunnan","admin","6250000","1156477539"
//"Nanchong","Nanchong","30.7991","106.0784","China","CN","CHN","Sichuan","minor","6183000","1156762337"
//"Zunyi","Zunyi","27.7050","106.9336","China","CN","CHN","Guizhou","","6127009","1156539782"
//"Lu’an","Lu'an","31.7542","116.5078","China","CN","CHN","Anhui","minor","6090000","1156499624"
//"Yichun","Yichun","27.8041","114.3830","China","CN","CHN","Jiangxi","minor","6048700","1156229590"
//
//