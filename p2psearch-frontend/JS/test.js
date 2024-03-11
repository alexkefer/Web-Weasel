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

// Create a text output element
const output = document.createElement('p');
output.setAttribute('id', 'output');

// Create a list element to display URLs
const urlListElement = document.createElement('ul');
urlListElement.setAttribute('id', 'urlList');

// Append all elements to the container
downloadContainer.appendChild(h1);
downloadContainer.appendChild(label);
downloadContainer.appendChild(input);
downloadContainer.appendChild(button);
downloadContainer.appendChild(output); // Append the text output
downloadContainer.appendChild(urlListElement); // Append the URL list

// Append the container to the main element
document.getElementById('main').appendChild(downloadContainer);

// Array to store URLs
const urlList = [];

// Function to display the URL list
function displayUrlList() {
    // Clear existing list
    //urlListElement.innerHTML = '';

    // Add each URL as a new list item with an active link
    urlList.forEach(url => {
        const listItem = document.createElement('li');
        const link = document.createElement('a');
        link.href = 'http://localhost:8080/retrieve?path=' + url;
        link.textContent = url;
        link.target = '_blank'; // Open link in a new tab
        listItem.appendChild(link);
        urlListElement.appendChild(listItem);
    });
}


// Add event listener to the button for fetch action
document.getElementById('fetchButton').addEventListener('click', function() {
    const urlInput = document.getElementById('urlInput').value.trim();
    
    // Regular expression to match URL pattern
    const urlRegex = /^(https?|http):\/\/[^\s$.?#].[^\s]*$/i;

    if (!urlInput || !urlRegex.test(urlInput)) {
        output.textContent = 'Please enter a valid URL.';
        return;
    }

    // Push the valid URL to the urlList array
    urlList.push(urlInput);

    // Update the display of the URL list
    displayUrlList();

    const url = 'http://localhost:8080/cache?path=' + encodeURIComponent(urlInput);

    fetch(url, { mode: 'no-cors' })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.text();
        })
        .then(data => {
            console.log('Cached URL:', data);
            output.textContent = 'Cached URL: ' + data; // Output the URL
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });
});

