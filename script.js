function handleImageUpload(inputId, ...previewIds) {
  const input = document.getElementById(inputId);

  input.addEventListener("change", function () {
    const file = this.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = function (e) {
        previewIds.forEach(previewId => {
          const preview = document.getElementById(previewId);
          if (preview) {
            preview.src = e.target.result;
          }
        });
      };
      reader.readAsDataURL(file);
    }
  });
}

// Handle all image uploads
handleImageUpload("profileUpload", "profilePreview", "profilePicPreview");
handleImageUpload("citizenshipFront", "frontPreview");
handleImageUpload("citizenshipBack", "backPreview");
