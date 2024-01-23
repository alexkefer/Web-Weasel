document.addEventListener('DOMContentLoaded', function () {
  // Get the iconButton element
  var iconButton = document.getElementById('iconButton');

  // Add a click event listener to the iconButton
  iconButton.addEventListener('click', function () {
    // Get the icon-img element
    var iconImg = document.querySelector('.icon-img');

    // Toggle between on_power_icon.png and off_power_icon.png based on the current src
    iconImg.src = (iconImg.src.includes('on_power_icon.png')) ? '../images/off_power_icon.png' : '../images/on_power_icon.png';
  });
});
