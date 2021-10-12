import React, { useState, useEffect } from 'react';

//Leaflet
import { MapContainer, TileLayer, Rectangle } from 'react-leaflet'

import quadtree from 'quadtree';

//Import Leaflet and plugins
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

//estado por defecto
const INITIAL_STATE = {
    lat: 43.2612,
    lng: -1.779,
    zoom: 8
}

const RenderMap = ({setAreaQuad}) => {
    const [mapState, setMapState] = useState(INITIAL_STATE); //Estado por defecto

    const [areaQuad, setAreaQuad] = useState(null); //Estado por defecto
    const [areaSelected, setAreaSelected] = useState(null); //Estado por defecto

    //Cambiar posición
    const changeLocation = e => setMapState({...mapState, lat: e.latlng.lat, lng: e.latlng.lng});

    //Corregir error de icono no encontrado
    //Si utilizas imagenes porpias para iconos o una libreria de marcadores, esto no hace falta
    useEffect(() => {
        delete L.Icon.Default.prototype._getIconUrl;
        L.Icon.Default.mergeOptions({
          iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
          iconUrl: require("leaflet/dist/images/marker-icon.png"),
          shadowUrl: require("leaflet/dist/images/marker-shadow.png")
        });
      }, []);


    //mouse over map
    const onMouseOver = (e) => {
      let location = e.latlng;
      let detail = mapState.zoom+1;
      let key = quadtree.encode(location, detail);
      let bbox = quadtree.bbox(key);
      setAreaQuad([[bbox.minlat, bbox.minlng],[bbox.maxlat, bbox.maxlng]]);
    }

    //mouse out map
    const onMouseOut = () => {
      setAreaQuad(null);
    }

    //on click area
    const onClickArea = (e) => {
      let location = e.latlng;
      let detail = mapState.zoom+1;
      let key = quadtree.encode(location, detail);
      let bbox = quadtree.bbox(key);
      setAreaSelected([[bbox.minlat, bbox.minlng],[bbox.maxlat, bbox.maxlng]]);
    }

    return (
        <MapContainer 
            center={[mapState.lat, mapState.lng]} 
            zoom={mapState.zoom} 
            style={{ height: 'calc(100vh - 210px)' }}
            onclick={ e => changeLocation(e)}

            //mouse over map and out map functions
            onmouseout={onMouseOut}
            onmousemove={onMouseOver}

            //get new zoom
            onzoomend= { d => setMapState(prevState => ({...prevState, zoom: d.target._zoom}))}
        >
          <TileLayer
            attribution='&amp;copy <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          />


          {
            //Area selected rectangle
            areaSelected && <Rectangle key={"rect2"} color="green" bounds={areaSelected}></Rectangle>
          }

          {
            //Area Over rectangle
            areaQuad && <Rectangle key={"rect"} color="blue" bounds={areaQuad} onclick={onClickArea}></Rectangle>
          }

        </MapContainer>
      )
};

export default RenderMap;


/*

PRUEBA DIFERENTES MAPAS SI QUIERES :)

const map_tiles = [
    {
        label: "Open Street Maps",
        value: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    },
    {
        label: "Satelite",
        value: "https://server.arcgisonline.com/ArcGIS/rest/services/World_Imagery/MapServer/tile/{z}/{y}/{x}",
        attribution: 'Tiles &copy; Esri &mdash; Source: Esri, i-cubed, USDA, USGS, AEX, GeoEye, Getmapping, Aerogrid, IGN, IGP, UPR-EGP, and the GIS User Community'
    },
    {
        label: "Open Street Maps DE",
        value: "https://{s}.tile.openstreetmap.de/tiles/osmde/{z}/{x}/{y}.png"
    },
    {
        label: "Open Street Maps HOT",
        value: "https://{s}.tile.openstreetmap.fr/hot/{z}/{x}/{y}.png"
    },
    {
        label: "Open Topo Map",
        value: "https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png"
    },
    {
        label: "Open Map Surfer",
        value: "https://maps.heigit.org/openmapsurfer/tiles/roads/webmercator/{z}/{x}/{y}.png"
    },
    {
        label: "Hydda Full",
        value: "https://{s}.tile.openstreetmap.se/hydda/full/{z}/{x}/{y}.png"
    },
    {
        label: "Esri World Street Map",
        value: "https://server.arcgisonline.com/ArcGIS/rest/services/World_Street_Map/MapServer/tile/{z}/{y}/{x}"
    },
    {
        label: "Esri World Topo Map",
        value: "https://server.arcgisonline.com/ArcGIS/rest/services/World_Topo_Map/MapServer/tile/{z}/{y}/{x}"
    }
]




*/