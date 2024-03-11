// Create a container div
const downloadContainer = document.createElement('div');
downloadContainer.setAttribute('id', 'downloadContainer');

// Create an h1 element
const h1 = document.createElement('h1');
h1.textContent = 'Fetch Example';

// Create a label element for the input field
const label = document.createElement('label');
label.textContent = 'Enter URL:';
label.setAttribute('for', 'urlInput');

// Create an input field
const input = document.createElement('input');
input.setAttribute('type', 'text');
input.setAttribute('id', 'urlInput');

// Create a button
const button = document.createElement('button');
button.setAttribute('id', 'fetchButton');
button.textContent = 'Fetch';

// Append all elements to the container
downloadContainer.appendChild(h1);
downloadContainer.appendChild(label);
downloadContainer.appendChild(input);
downloadContainer.appendChild(button);

// Append the container to the main element
document.getElementById('main').appendChild(downloadContainer);

// Add event listener to the button for fetch action
document.getElementById('fetchButton').addEventListener('click', function() {
    const urlInput = document.getElementById('urlInput').value.trim();
    if (!urlInput) {
        alert('Please enter a valid URL.');
        return;
    }

    const url = 'http://localhost:8080/cache?path=' + encodeURIComponent(urlInput);

    fetch(url)
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.text();
        })
        .then(data => {
            console.log('Cached URL:', data);
            // Further processing of the cached URL
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });
});
