// Wait for the DOM content to be fully loaded before executing the script
document.addEventListener("DOMContentLoaded", function() {
    // Get the sidebar element by its ID
    var sidebar = document.getElementById("sidebar");

    // Array containing the sidebar items with their names and URLs
    var sidebarItems = [
        { name: "Home", url: "index.html" },
        { name: "Tutorial", url: "tutorial.html" },
        { name: "Caching", url: "caching.html" },
        { name: "Settings", url: "settings.html" },
        { name: "Resources", url: "resources.html" }
    ];

    // Loop through each item in the sidebarItems array
    sidebarItems.forEach(function(item) {
        // Create an 'a' element for each sidebar item
        var link = document.createElement("a");

        // Set the 'href' attribute of the 'a' element to the URL of the sidebar item
        link.href = item.url;

        // Set the text content of the 'a' element to the name of the sidebar item
        link.textContent = item.name;

        // Add a CSS class to the 'a' element for styling purposes
        link.classList.add("sidebar-item");

        // Append the 'a' element to the sidebar
        sidebar.appendChild(link);
    });
});
