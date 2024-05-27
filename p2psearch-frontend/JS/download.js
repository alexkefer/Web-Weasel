// Retrieve stored URL list from the server on page load
window.addEventListener('load', function() {
    fetch('http://localhost:8080/sites')
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.text(); // Get the response as plain text
        })
        .then(text => {
            const storedUrlList = text.split('\n').filter(url => url.trim() !== '');
            if (storedUrlList.length > 0) {
                urlList.push(...storedUrlList);
                displayUrlList();
            }
        })
        .catch(error => {
            console.error('There was a problem with fetching the URL list:', error);
        });
});

// Function to clear a specific URL from the list and update the display
function clearUrl(url) {
    const trimmedUrl = url.replace(/^https?:\/\//, '');
    fetch('http://localhost:8080/removeSite?url=' + trimmedUrl, {
        method: 'GET', // or 'DELETE' if supported by the server
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        // Remove the URL from the local list and update the display
        const index = urlList.indexOf(url);
        if (index !== -1) {
            urlList.splice(index, 1);
            displayUrlList();
            // Display output message
            output.textContent = 'URL successfully removed: ' + url;
        }
    })
    .catch(error => {
        console.error('There was a problem with removing the URL:', error);
    });
}


// Array to store URLs
const urlList = [];

// Function to display the URL list
function displayUrlList() {
    // Clear existing list
    urlListElement.innerHTML = '';

    // Add each URL as a new list item with an active link
    urlList.forEach(url => {
        // Trim off "https://" from the URL
        const trimmedUrl = url.replace(/^https?:\/\//, '');

        const listItem = document.createElement('li');
        const link = document.createElement('a');
        link.href = 'http://localhost:8080/retrieve?path=' + trimmedUrl;
        link.textContent = trimmedUrl; // Display trimmed URL
        link.target = '_blank'; // Open link in a new tab

        // Create a button to clear this URL
        const clearButton = document.createElement('button');
        clearButton.textContent = 'Clear';
        clearButton.addEventListener('click', function() {
            clearUrl(url);
        });

        // Append the link and clear button to the list item
        listItem.appendChild(link);
        listItem.appendChild(clearButton);

        // Append the list item to the URL list
        urlListElement.appendChild(listItem);
    });
}


// Create the necessary HTML elements
const downloadContainer = document.createElement('div');
downloadContainer.setAttribute('id', 'downloadContainer');

const h1 = document.createElement('h1');
h1.textContent = 'Web Caching';

const label = document.createElement('label');
label.textContent = 'Enter URL:';
label.setAttribute('for', 'urlInput');

const input = document.createElement('input');
input.setAttribute('type', 'text');
input.setAttribute('id', 'urlInput');

const button = document.createElement('button');
button.setAttribute('id', 'fetchButton');
button.textContent = 'Fetch';

const output = document.createElement('p');
output.setAttribute('id', 'output');

const urlListTitle = document.createElement('h2');
urlListTitle.textContent = 'Saved URLs';

const urlListElement = document.createElement('ul');
urlListElement.setAttribute('id', 'urlList');

// Append all elements to the container
downloadContainer.appendChild(h1);
downloadContainer.appendChild(label);
downloadContainer.appendChild(input);
downloadContainer.appendChild(button);
downloadContainer.appendChild(output); // Append the text output
downloadContainer.appendChild(urlListTitle);
downloadContainer.appendChild(urlListElement); // Append the URL list

// Append the container to the main element
document.getElementById('main').appendChild(downloadContainer);

// Add event listener to the button for fetch action
document.getElementById('fetchButton').addEventListener('click', function() {
    let urlInput = document.getElementById('urlInput').value.trim();
    
    // Regular expression to match URL pattern
    const urlRegex = /^(https?|http):\/\/[^\s$.?#].[^\s]*$/i;

    // Automatically prepend "https://" if missing
    if (!urlInput.startsWith('http://') && !urlInput.startsWith('https://')) {
        urlInput = 'https://' + urlInput;
    }

    // Validate the URL
    if (!urlRegex.test(urlInput)) {
        output.textContent = 'Please enter a valid URL.';
        return;
    }

    // Check if the URL is already in the list
    if (urlList.includes(urlInput)) {
        output.textContent = 'URL is already in the list.';
        return;
    }

    // Push the valid URL to the urlList array
    urlList.push(urlInput);

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

            // Display the URL list after showing the cached URL
            displayUrlList();
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });
});
