import React from 'react';

function UploadComponent() {
    const handleFileUpload = (event) => {
        // Handle file upload logic here
    };

    return (
        <div>
            <input type="file" onChange={handleFileUpload} />
        </div>
    );
}

export default UploadComponent;
