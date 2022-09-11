<script>
  import { storeFE, groceryList } from '../scripts/store.js'

  import {
    Modal,
    ModalFooter,
    Container,
    Row
  } from 'sveltestrap'

  let open = false;
  // const toggle = () => (open = !open);
  function toggle() {
    open = !open;
    getIngredients();
  }

  let groceries = [];

  async function getIngredients() {
    let response = await fetch("http://127.0.0.1:8000/ingredients", {
			method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
			body: JSON.stringify($storeFE),
		});
    
    let ingredients = await response.json();
    groceries = ingredients;
  }

  function printIngredient(ingredient) {
    if (ingredient.quantity == "" && ingredient.measurement == "") {
      return ingredient.item;
    }
    if (ingredient.quantity != "" && ingredient.measurement != ""){
      return ingredient.quantity + " " + ingredient.measurement + " " + ingredient.item;
    }
    if (ingredient.quantity != "" && ingredient.measurement == "") {
      return ingredient.quantity + " " + ingredient.item;
    }
    
  }
</script>

<div>
  <button class="button button_normal" on:click={toggle} disabled={$storeFE.length == 0}>Grocery List</button>
  <Modal isOpen={open} {toggle}>
    <div class="modal_header">
      <h1>Grocery List</h1>
    </div>
    <div class="modal_body">
      <Container>
        <Row cols={1} style="max-height:calc(100vh - 300px); overflow:scroll">
          {#each groceries as ingredient}
            <div class="checkbox">
              <input id="ingredient" type ="checkbox" >
              <label for="ingredient"> {printIngredient(ingredient)} </label>
            </div>
          {/each}
        </Row>
      </Container>
    </div>
  </Modal>
</div>

<style>
	.button_normal {
		background:#05386B; 
		color:#EDF5E1;
		border-radius: 8px;
	}

  .button_normal:disabled {
		background-color: darkslategrey;
	}

  .checkbox input,
  .checkbox label {
    color: #05386B;
    /* width: auto; */
    margin-left: 0;
    display: inline-block;
    vertical-align: middle;
    font-family: Georgia, 'Times New Roman', Times, serif;
    /* float: left; */
  }

  .checkbox:checked {
    background-color: #05386B;
    margin: 20px;
  }

  .modal_header {
    background:#379683; 
    color:#EDF5E1;
    display: block;
  }

  .modal_body {
    background:#EDF5E1;
    color: #05386B;
    padding: 10px;
  }

  h1 {
    margin: 20px;
    font-family: Georgia, 'Times New Roman', Times, serif;
  }

</style>
