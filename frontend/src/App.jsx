import React, { useState } from "react";
import CertificateForm from "./components/CertificateForm";
import CertificatePreview from "./components/CertificatePreview";
import { generateCertificate } from "./api";

function App() {
    const [formData, setFormData] = useState({
        name: "",
        company: "",
        position: "",
        duration: "",
        start_date: "",
        end_date: "",
        template: ""
    });

    const templates = [
        {
            id: "prodigy",
            name: "Prodigy Infotech",
            previewImage: "prodigy.png",
            fields: {
                name: { x: 200, y: 150, font_size: 24, font_family: "Arial", color: [0, 0, 0] },
                download_date: { x: 450, y: 300, font_size: 14, font_family: "Arial", color: [0, 0, 0] }
            }
        }
    ];

    const [selectedTemplate, setSelectedTemplate] = useState(null);

    const handleTemplateSelect = (templateId) => {
        const tpl = templates.find((t) => t.id === templateId);
        setSelectedTemplate(tpl);
        setFormData({ ...formData, template: templateId });
    };

    const handleDownload = async () => {
        try {
            const pdfBlob = await generateCertificate(formData);
            const url = window.URL.createObjectURL(new Blob([pdfBlob]));
            const a = document.createElement("a");
            a.href = url;
            a.download = "certificate.pdf";
            a.click();
        } catch (err) {
            console.error("Error generating certificate:", err);
        }
    };

    const handleReset = () => {
        setFormData({
            name: "",
            company: "",
            position: "",
            duration: "",
            start_date: "",
            end_date: "",
            template: ""
        });
        setSelectedTemplate(null);
    };

    return (
        <div className="app-container">
            <CertificateForm
                formData={formData}
                setFormData={setFormData}
                templates={templates}
                onTemplateSelect={handleTemplateSelect}
                onSubmit={handleDownload}
                onReset={handleReset}
            />
            <CertificatePreview formData={formData} selectedTemplate={selectedTemplate} />
        </div>
    );
}

export default App;
