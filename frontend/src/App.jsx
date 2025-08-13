import React, { useState } from "react";
import CertificateForm from "./components/CertificateForm";
// import CertificatePreview from "./components/CertificatePreview";
import "./App.css";

export default function App() {
  const [previewData, setPreviewData] = useState(null);
  console.log('previewData:', previewData);

  return (
    <div className="app">
      <h1>Certificate Generator</h1>
      <div className="main-grid">
        <CertificateForm  />
        {/* <CertificatePreview formData={previewData} setFormData={setPreviewData} /> */}
      </div>
    </div>
  );
}

