import React from "react";

export default function CertificatePreview({ formData, selectedTemplate }) {
    if (!selectedTemplate) return <div className="preview-placeholder">Select a template to preview</div>;

    return (
        <div className="certificate-preview" style={{ position: "relative", display: "inline-block" }}>
            <img
                src={`/templates/${selectedTemplate.previewImage}`}
                alt="Template Preview"
                style={{ width: "100%", border: "1px solid #ccc" }}
            />
            {/* Name */}
            <div style={{
                position: "absolute",
                top: `${selectedTemplate.fields.name.y}px`,
                left: `${selectedTemplate.fields.name.x}px`,
                fontSize: selectedTemplate.fields.name.font_size,
                fontFamily: selectedTemplate.fields.name.font_family,
                color: `rgb(${selectedTemplate.fields.name.color.join(",")})`
            }}>
                {formData.name}
            </div>
            {/* Date */}
            <div style={{
                position: "absolute",
                top: `${selectedTemplate.fields.download_date.y}px`,
                left: `${selectedTemplate.fields.download_date.x}px`,
                fontSize: selectedTemplate.fields.download_date.font_size,
                fontFamily: selectedTemplate.fields.download_date.font_family,
                color: `rgb(${selectedTemplate.fields.download_date.color.join(",")})`
            }}>
                {new Date().toLocaleDateString("en-GB")}
            </div>
        </div>
    );
}
