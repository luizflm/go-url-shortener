const urlInput = document.querySelector('#url');
const commandUrl = document.querySelector('#command-url');

if (urlInput && commandUrl) {
  const updateCommand = () => {
    commandUrl.textContent = urlInput.value;
  };

  urlInput.addEventListener('input', updateCommand);
  updateCommand();
}
