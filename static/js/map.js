let leafletMap = document.querySelector('#mapid');

if (leafletMap) {

    var mymap = L.map('mapid').setView([56.94965, 24.1052], 13);

    var CustomIcon = L.Icon.extend({
        options: {
            iconSize:     [40, 40],
            iconAnchor:   [20, 40],
            popupAnchor:  [0, -80]
        }
    });
    
    mapLink = '<a href="http://openstreetmap.org">OpenStreetMap</a>';

    L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
        attribution: 'Map data &copy; ' + mapLink + ' contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
        maxZoom: 24,
        id: 'mapbox.streets',
        accessToken: 'pk.eyJ1Ijoib3N3ZWUiLCJhIjoiY2ppMXdycTZyMGY3ZDNwcXBzZnUzcTBiYiJ9.mi4BGXgJbdDgS7mVGVyExA',
        style: 'mapbox://styles/mapbox/basic-v9',
    }).addTo(mymap);

    for (var i = 0; i < markers.length; i++) {
        svgicon = `<svg id="Marker" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
        <defs>
          <style>
            .cls-1 {
              fill: #006aff;
            }
      
            .cls-2 {
              font-size: 28px;
              fill: #fff;
              font-family: Montserrat-Regular, Montserrat;
              text-align: center;
            }
          </style>
        </defs>
        <title>marker-01</title>
        <path class="cls-1" d="M56.9679,24.9679A24.9679,24.9679,0,1,0,12.38928,40.42236h-.00067l18.4848,22.96552a1.38676,1.38676,0,0,0,2.22162.10583l17.86951-22.285H50.963A24.86853,24.86853,0,0,0,56.9679,24.9679Z"/>
        <text class="cls-2" transform="translate(18 34)">` + markers[i][1] + `</text>
      </svg>`
    
        url = encodeURI("data:image/svg+xml," + svgicon).replace('#','%23');
    
        icon = new CustomIcon({iconUrl: url})
        marker = new L.marker([markers[i][2],markers[i][3]], {icon: icon})
            .bindPopup(markers[i][0])
            .addTo(mymap);
    };
}