let leafletMap = document.querySelector('#mapid');

var geojson =  [{
"type": "Feature",
"geometry": {
"type": "Point",
"coordinates": [-106.62987, 35.10418]
 },
 "properties": {
 "name": "My Point",
 "title": "A point at the Big I"
 }
},{
"type": "Feature",
 "geometry": {
 "type": "LineString",
 "coordinates":[[-106.67999, 35.14097],
 [-106.68892, 35.12974],
 [-106.69064, 35.1098]]
 },
"properties": {
 "name": "My LineString",
 "title": "A line along the Rio Grande"
 }
},{
"type": "Feature",
 "geometry": {
 "type": "Polygon",
 "coordinates":[[[-106.78059, 35.14574],
 [-106.7799, 35.10559],
 [-106.71467, 35.13704],
 [-106.69716, 35.17942],
 [-106.78059, 35.14574]]]
 },
"properties": {
 "name": "My Polygon",
 "title": "Balloon Fiesta Park"
 }
 }];

if (leafletMap) {
    // function foundLocation(e){
    //     var mydate = new Date(e.timestamp);
    //     L.marker(e.latlng).addTo(mymap).bindPopup(mydate.toString());
    // }
    // function notFoundLocation(e){
    //     alert("Unable to get your location for centering map to it. You may need to enable Geolocation.");
    // }

    var map = L.map('mapid').setView([56.94965, 24.1052], 13);

    // mymap.on('locationfound', foundLocation);
    // mymap.on('locationerror', notFoundLocation);

    // mymap.locate({setView: true, maxZoom:13});

    ////mymap.on('click', function(){alert("You clicked the map");});
    // mymap.on('click',function(e){
    //     var coord=e.latlng.toString().split(',');
    //     var lat=coord[0].split('(');
    //     var long=coord[1].split(')');
    //     L.marker(e.latlng).addTo(mymap);
    // });

    // var myMarker = L.marker([56.94965, 24.1052],{title:"MyPoint", alt:"The Big I", draggable:true}).addTo(mymap);;
    // function whereAmI(){myMarker.bindPopup("I have been moved to:"+ String(myMarker.getLatLng()))};
    // myMarker.on('dragend', whereAmI);

    // var marker1 = L.marker([56.94965, 24.2052]).addTo(mymap).bindPopup(createPopup("Text as a parameter"));
    // var marker2 = L.marker([56.94965, 24.0]).addTo(mymap).bindPopup(createPopup("Different text as a parameter"));

    // function createPopup(x){
    //     return L.popup({keepInView:true, closeButton:false}).setContent(x);
    // }



    function styleFunction(){return {color: "purple"}; }
    function newStyle(){geoJsonLayer.setStyle({color:"green"});}
    function oldStyle(){geoJsonLayer.setStyle({color:"purple"});}
    
    var geoJsonLayer = L.geoJson(geojson, {
        style:styleFunction,
        onEachFeature: function(feature,layer) {
            layer.bindPopup(feature.properties.title);
        }
    }).addTo(map);
    
    geoJsonLayer.on('mouseover',newStyle);geoJsonLayer.on('mouseout',oldStyle);

    // var points = [
    //     [35.1555 , -106.591838 , "<img src='http://farm8.staticflickr.com/7153/6831137393_fa38634fd7_m.jpg'>"],
    //     [35.0931 , -106.664177 , "<img src='http://farm3.staticflickr.com/2167/2479129916_0d861b2600.jpg'>"],
    //     [35.1143 , -106.577991 , "<img src='http://farm2.staticflickr.com/1416/908720823_e390a242f4.jpg'>"]
    // ];

    // var heat = L.HeatLayer(points).addTo(map);




    
    var CustomIcon = L.Icon.extend({
        options: {
            iconSize:     [40, 40],
            iconAnchor:   [20, 40],
            popupAnchor:  [0, -80]
        }
    });
    
    mapLink = '<a href="http://openstreetmap.org">OpenStreetMap</a>';

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: 'Map data &copy; ' + mapLink + ' contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>',
        maxZoom: 24,
        //id: 'mapbox.streets',
        //accessToken: 'pk.eyJ1Ijoib3N3ZWUiLCJhIjoiY2ppMXdycTZyMGY3ZDNwcXBzZnUzcTBiYiJ9.mi4BGXgJbdDgS7mVGVyExA',
        //style: 'mapbox://styles/mapbox/basic-v9',
    }).addTo(map);

    //var polygon = L.polygon([[56.95898, 24.013582],[56.971394, 24.137053],[56.956596, 24.156408],[56.934544, 24.159397], [56.941505, 24.11933]], {color: '#006aff', weight:3, fillColor:'#FACE20',fillOpacity:0.5}).addTo(mymap);
    for (var i = 0; i < markers.length; i++) {
        svgicon = `<svg id="marker` + markers[i][4] + `" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
        <defs>
          <style>
            .cls-1 {
              fill: #006aff;
            }      
            .cls-2 {
              font-size: 28px;
              fill: #fff;
              font-family: Montserrat-Regular, Montserrat;
            }
          </style>
        </defs>
        <path class="cls-1" d="M56.9679,24.9679A24.9679,24.9679,0,1,0,12.38928,40.42236h-.00067l18.4848,22.96552a1.38676,1.38676,0,0,0,2.22162.10583l17.86951-22.285H50.963A24.86853,24.86853,0,0,0,56.9679,24.9679Z"/>
        <text id="label` + markers[i][4] + `" class="cls-2" transform="translate(32 34)" text-anchor="middle">` + markers[i][1] + `</text>
      </svg>`
    
        //url = encodeURI("data:image/svg+xml," + svgicon).replace('#','%23');
        url = encodeURI("data:image/svg+xml;utf8," + svgicon);
    
        icon = new CustomIcon({iconUrl: url})

        marker = new L.marker([markers[i][2], markers[i][3]], options={
            //"title" : "Apollo",
            icon: icon,
        })
        //.bindPopup(markers[i][0] + ' ID: ' + markers[i][4])
        .bindTooltip(`<p>` + markers[i][0] + `</p>`)
        .addTo(map);
    };
}