package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var myLat  float64
var myLong float64
var myDist int
var myDebug bool

type feed struct {
	Id        int    `json: "id"`
	Name      string `json: "name"`
	PolarPlot bool   `json: "polarPlot"`
}

var planeType = map[int]string {
	0: "None",
	1: "Land Plane",
	2: "Sea Plane",
	3: "Amphibian",
	4: "Helecopter",
	5: "Gyrocopter",
	6: "Tiltwing",
	7: "Groud Vehicle",
	8: "Tower",
}

type info struct {
	Feeds []feed
	AcList []ac
	Src       int    `json: "src"`
	SrcFeed   int    `json: "srcFeed"`
	ShowSil   bool   `json: "showSil"`
	ShowFlg   bool   `json: "showFlg"`
	ShowPic   bool   `json: "showPic"`
	FlgH      int    `json: "flgH"`
	FlgW      int    `json: "flgW"`
	TotalAc   int    `json: "totalAc"`
	LastDv    string `json: "lastDv"`
	ShtTrlSec int    `json: "shtTrlSec"`
	Stm       int64  `json: "stm"`
}

type ac struct {
	Id           int    `json: "Id"`
	Rcvr         int    `json: "Rcvr"`
	HasSig       bool   `json: "HasSig"`
	Sig          int    `json: "Sig"`
	Icao         string `json: "Icao"`
	Bad          bool   `json: "Bad"`
	Reg          string `json: "Reg"`
	FSeen        string `json: "FSeen"`
	TSecs        int    `json: "TSecs"`
	CMsgs        int    `json: "CMsgs"`
	Alt          int    `json: "Alt"`
	GAlt         int    `json: "GAlt"`
	InHG         float32 `json: "InHG"`
	AltT         int    `json: "AltT"`
	Call         string `json: "Call"`
	Lat          float32    `json: "Lat"`
	Long         float32    `json: "Long"`
	PosTime      int64    `json: "PosTime"`
	Mlat         bool   `json: "Mlat"`
	Tisb         bool   `json: "Tisb"`
	Spd          float32    `json: "Spd"`
	Trak         float32    `json: "Trak"`
	TrkH         bool   `json: "TrkH"`
	Type         string `json: "Type"`
	Mdl          string `json: "Mdl"`
	Man          string `json: "Man"`
	CNum         string `json: "CNum"`
	Op           string `json: "Op"`
	OpIcao       string `json: "OpIcao"`
	From         string `json: "From"`
	To           string `json: "To"`
	Stops        []string `json: "Stops"`
	Sqk          string `json: "Sqk"`
	Help         bool   `json: "Help"`
	Vsi          int    `json: "Vsi"`
	VsiT         int    `json: "VsiT"`
	Dst          float32    `json: "Dst"`
	Brng         float32    `json: "Brng"`
	WTC          int    `json: "WTC"`
	Species      int    `json: "Species"`
	EngType      int    `json: "EngType"`
	EngMount     int    `json: "EngMount"`
	Mil          bool   `json: "Mil"`
	Cou          string `json: "Cou"`
	HasPic       bool   `json: "HasPic"`
	Interested   bool   `json: "Interested"`
	FlightsCount int    `json: "FlightsCount"`
	Gnd          bool   `json: "Gnd"`
	SpdType      int    `json: "SpdType"`
	CallSus      bool   `json: "CallSus"`
	Trt          int    `json: "Trt"`
	Year         string `json: "Year"`
}

func main() {
	flag.Float64Var(&myLat, "lat",29.4608,"my latatude")
	flag.Float64Var(&myLong, "long",-95.0513,"my latatude")
	flag.IntVar(&myDist, "dist", 20, "search distance")
	flag.BoolVar(&myDebug, "debug", false, "print extra debug output")
	flag.Parse()

	fmt.Printf("Searching for air planes within %d km of (%f, %f)\n", myDist, myLat, myLong)
	res, err := http.Get("http://public-api.adsbexchange.com/VirtualRadar/AircraftList.json?lat="+ fmt.Sprintf("%f",myLat) +"&lng="+ fmt.Sprintf("%f",myLong)+ "&fDstL=0&fDstU="+ fmt.Sprintf("%d",myDist) )

	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if json.Valid([]byte(jsonData)) == true {
		if myDebug == true {
			fmt.Printf("Valid JSON object returned\n")
			fmt.Printf("%s", jsonData)
			fmt.Println("")
		}
	}

	flightData := info{}

	err = json.Unmarshal([]byte(jsonData), &flightData )

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf( "Number of air plans in area: %d\n" , len(flightData.AcList))
	for index, plane  := range flightData.AcList {
		fmt.Printf("Plane %d distance is %f at altatude %d going %f\n",index, plane.Dst, plane.Alt, plane.Spd )
		fmt.Printf(" Call sign %s. Registration %s\n", plane.Call, plane.Reg )
		fmt.Printf(" Manufacture: %s\n", plane.Man )
		fmt.Printf(" Model: %s\n", plane.Mdl )
		fmt.Printf(" Type: %s\n", plane.Type )
		fmt.Printf(" Airplane type: %s\n", planeType[plane.Species] )
		if plane.Op != "" {
			fmt.Printf(" Operated by %s\n", plane.Op )
		}
		if plane.From != "" {
			fmt.Printf(" Flying from %s to %s\n", plane.From, plane.To)
		}
		if plane.Mil == true {
			fmt.Printf(" It is a military plane\n")
		}
		if len(plane.Stops) > 0 {
			fmt.Printf("  with stops in " )
			for stopNo, stop := range plane.Stops {
				fmt.Printf("%s(%d), ", stop, stopNo )
			}
			fmt.Println("")
		}
	}

}

