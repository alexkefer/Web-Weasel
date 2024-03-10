const mockFolders = [
    { name: 'My Resource', files: ['file1.txt', 'file2.txt'] },
    { name: 'Peer 1', files: ['file3.txt', 'file4.txt'] },
    { name: 'Peer 2', files: ['file5.txt', 'file6.txt'] },
    { name: 'Peer 3', files: ['file7.txt', 'file8.txt'] }
];

displayFolderList(mockFolders);

function displayFolderList(folders) {
    const fileListElement = document.getElementById('fileList');

    folders.forEach(folder => {
        const folderItem = document.createElement('li');
        folderItem.className = 'folder-list-item';

        const folderToggle = document.createElement('span');
        folderToggle.className = 'folder-toggle';
        folderToggle.textContent = '▶ '; // Default: right-pointing arrow
        folderToggle.addEventListener('click', () => toggleFolder(folderItem, folderToggle));

        const folderText = document.createTextNode(folder.name);
        folderItem.appendChild(folderToggle);
        folderItem.appendChild(folderText);

        const subFileList = document.createElement('ul');
        subFileList.className = 'file-list'; // Corrected class name

        folder.files.forEach(file => {
            const listItem = document.createElement('li');
            const downloadButton = document.createElement('button');
            downloadButton.className = 'download-button';
            downloadButton.textContent = `Download ${file}`;
            downloadButton.addEventListener('click', () => downloadFile(file));

            listItem.appendChild(downloadButton);
            subFileList.appendChild(listItem);
        });

        folderItem.appendChild(subFileList);
        fileListElement.appendChild(folderItem);
    });
}

function downloadFile(filename) {
    // Create a sample content for the file
    const sampleContent = 'This is a sample content for ' + filename;

    // Create a Blob with the sample content
    const blob = new Blob([sampleContent], { type: 'text/plain' });

    // Create a download link
    const downloadLink = document.createElement('a');
    downloadLink.href = URL.createObjectURL(blob);
    downloadLink.download = filename;

    // Append the link to the document
    document.body.appendChild(downloadLink);

    // Trigger a click on the link to start the download
    downloadLink.click();

    // Remove the link from the document
    document.body.removeChild(downloadLink);
}
