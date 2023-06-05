/*
isAdd : bool, true if add, false if sub
id : id of the item to change
elm : html element that called the function
*/
function stockpileItemNumberClick(isAdd, id) {
   // Send the AJAX request to the backend
   console.log("Sending to stockpile :", isAdd, id)
  fetch('/api/stockpileUpdateItem', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    // JSON format send here
    body: JSON.stringify({ isAdd: isAdd === "true", id: id }),
  })
    .then(response => response.json())
    .then(data => {
        
        console.log(data);
      // Handle the response from the backend
      if (data.status === 'success') {

        // LOGIC GOES HERE
        var numberElement = document.getElementById("nb#"+id); ;
        var itemBoxElement = document.getElementById("itemBox#"+id);
        numberElement.textContent = data.itemNumber;

        // remove all border class
        itemBoxElement.classList.remove("borderCritical", "borderLow", "borderNotFull", "borderFull", "borderOverfilled")
        // set class
        itemBoxElement.classList.add(data.class);

      } else {
        console.error('Failed to update item count:', data.error);
      }
    })
    .catch(error => {
      console.error('An error occurred while updating item count:', error);
    });
  }
