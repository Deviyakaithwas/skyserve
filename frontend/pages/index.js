import React from 'react';
import MapComponent from '../components/MapComponent';
import UploadComponent from '../components/UploadComponent';

function Home() {
    return (
        <div>
            <h1>Geospatial Data Management</h1>
            <UploadComponent />
            <MapComponent />
        </div>
    );
}

export default Home;
