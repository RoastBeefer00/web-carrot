<script>
  import { storeFE, groceryList } from '../scripts/store.js'
  import { derived } from 'svelte/store';

  import {
    Button,
    Modal,
    ModalHeader,
    ModalFooter,
    ModalBody,
    Input,
    Container,
    Row
  } from 'sveltestrap'

  let open = false;
  // const toggle = () => (open = !open);
  function toggle() {
    open = !open;
    splitIngredients();
  }

  $: sortedGroceryList = sortByIngredient(Object.values($groceryList));
  $: combinedGroceryList = combineIngredients(sortedGroceryList);

  function sortByIngredient(array) {
    array.sort((el1, el2) =>
      el1.ingredient.toLowerCase().localeCompare(el2.ingredient.toLowerCase())
    );
    return array;
  }

  function splitIngredients() {
    $groceryList = [];
    var value;
    var measurement;
    var rest;
    var json;
    $storeFE.forEach(recipe => {
      recipe.ingredients.forEach(ingredient => {
      value = ingredient.match("^\\d*[^a-zA-Z \\*]?\\d*|¼|¾|½|⅓");
      console.log(value);
      measurement = ingredient.match("tbsps?|tsps?|cups?|cans?|packages?|packets?|ozs?|pounds?");
      console.log(measurement);
      
      if (measurement == null) {
        rest = ingredient.match("\\*?[a-zA-Z].*");
        console.log(rest);
      }
      else{
        rest = ingredient.match("(?<=tbsps? |tsps? |cups? |cans? |packages? |packets? |ozs? |pounds? )\\*?[a-zA-Z].*");
        console.log(rest);
      }

      json = "{\"quantity\":\"" + value + "\", \"measurement\":\"" + measurement + "\", \"ingredient\":\"" + rest + "\"}";
      $groceryList[$groceryList.length] = JSON.parse(json);
      });
    });
  }

  function combineIngredients(array) {
    for (let i = 0; i < array.length - 1; i++) {
      var j = i + 1;
      if (array[i].ingredient.includes(array[j].ingredient) || array[j].ingredient.includes(array[i].ingredient)) {
        if (array[i].measurement.includes(array[j].measurement) || array[j].measurement.includes(array[i].measurement)) {
          if (array[i].quantity != "") {
            if (array[i].ingredient.length > array[j].ingredient.length){
              array[i].ingredient = array[j].ingredient;
            }
            if (array[j].measurement.length > array[i].measurement.length){
              array[i].measurement = array[j].measurement;
            }
            array[i].quantity = parseFloat(array[i].quantity) + parseFloat(array[j].quantity);
          }
          array.splice(j, 1);
          i--;
        }
      }
    }
    return array;
  }

  function printIngredient(ingredient) {
    if (ingredient.quantity == '' && ingredient.measurement == "null") {
      return ingredient.ingredient;
    }
    if (ingredient.quantity != '' && ingredient.measurement != "null"){
      return ingredient.quantity + " " + ingredient.measurement + " " + ingredient.ingredient;
    }
    if (ingredient.quantity != '' && ingredient.measurement == "null") {
      return ingredient.quantity + " " + ingredient.ingredient;
    }
    
  }
</script>

<div>
  <Button style="background:#05386B; color:#EDF5E1" on:click={toggle} disabled={$storeFE.length == 0}>Grocery List</Button>
  <Modal isOpen={open} {toggle}>
    <ModalHeader {toggle} style="background:#379683; color:#EDF5E1">Grocery List</ModalHeader>
    <ModalBody style="background:#EDF5E1">
      <Container>
        <Row cols={1} style="max-height:calc(100vh - 300px); overflow:scroll">
          {#each combinedGroceryList as ingredient}
              <Input style="color:#05386B" id="c1" type ="checkbox" label={printIngredient(ingredient)} />

          {/each}
        </Row>
      </Container>
    </ModalBody>
    <ModalFooter style="background:#379683">
      <!-- <Button color="secondary" on:click={splitIngredients} style="background:#05386B">Close</Button> -->
    </ModalFooter>
  </Modal>
</div>