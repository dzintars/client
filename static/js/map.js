let leafletMap = document.querySelector('#mapid');

if (leafletMap) {
    // var markers = [
    //     ["Rīga", 56.94965, 24.1052],
    //     ["Oswee", 56.96, 24.14],
    //     ["Oswee X", 56.98, 24.16],
    // ];
    var mymap = L.map('mapid').setView([56.94965, 24.1052], 13);

    var CustomIcon = L.Icon.extend({
        options: {
            iconSize:     [40, 40],
            shadowSize:   [50, 64],
            iconAnchor:   [20, 40],
            shadowAnchor: [0, 40],
            popupAnchor:  [0, -80]
        }
    });

    var svgicon = "<svg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' x='0px' y='0px' viewBox='0 0 365 560' enable-background='new 0 0 365 560' xml:space='preserve'><style type='text/css'>.st0{fill:#00AEEF;} .st1{font-family:'Montserrat-Light';} .st2{font-size:256px;}</style><g><path fill='#006AFF' d='M182.9,551.7c0,0.1,0.2,0.3,0.2,0.3S358.3,283,358.3,194.6c0-130.1-88.8-186.7-175.4-186.9 C96.3,7.9,7.5,64.5,7.5,194.6c0,88.4,175.3,357.4,175.3,357.4S182.9,551.7,182.9,551.7z M122.2,187.2c0-33.6,27.2-60.8,60.8-60.8 c33.6,0,60.8,27.2,60.8,60.8S216.5,248,182.9,248C149.4,248,122.2,220.8,122.2,187.2z'/><text transform='matrix(1 0 0 1 169 207)' class='st1 st2'>1</text></g></svg>";

    var url = encodeURI("data:image/svg+xml," + svgicon).replace('#','%23');

    var icon = new CustomIcon({iconUrl: url})

    // var LeafIcon = L.Icon.extend({
    //     options: {
    //         shadowUrl: 'static/img/png/leaf-shadow.png',
    //         iconSize:     [38, 95],
    //         shadowSize:   [50, 64],
    //         iconAnchor:   [22, 94],
    //         shadowAnchor: [4, 62],
    //         popupAnchor:  [-3, -76]
    //     }
    // });

    // var greenIcon = new LeafIcon({iconUrl: 'static/img/png/leaf-green.png'}),
    //     redIcon = new LeafIcon({iconUrl: 'static/img/png/leaf-red.png'}),
    //     orangeIcon = new LeafIcon({iconUrl: 'static/img/png/leaf-orange.png'});

    // var marker = L.marker([56.94965, 24.1052]).addTo(mymap);
    mapLink = '<a href="http://openstreetmap.org">OpenStreetMap</a>';
    L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
        attribution: 'Map data &copy; ' + mapLink + ' contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
        maxZoom: 24,
        id: 'mapbox.streets',
        accessToken: 'pk.eyJ1Ijoib3N3ZWUiLCJhIjoiY2ppMXdycTZyMGY3ZDNwcXBzZnUzcTBiYiJ9.mi4BGXgJbdDgS7mVGVyExA',
        style: 'mapbox://styles/mapbox/basic-v9',
    }).addTo(mymap);

    for (var i = 0; i < markers.length; i++) {
        marker = new L.marker([markers[i][2],markers[i][3]], {icon: icon})
            .bindPopup(markers[i][0])
            .addTo(mymap);
    };
}