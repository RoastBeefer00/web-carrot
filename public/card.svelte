<script>
    async function removeRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Delete",
                    "index": index,
                    "recipe": recipe
                }
            } 
        }

        let response = await fetch("http://127.0.0.1:8000/delete", {
			method: 'DELETE',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
			body: JSON.stringify({
                recipe: recipe,
                store: $storeFE
            }),
		});
        let recipes = await response.json();
        $storeFE = recipes;
	}

    async function replaceRecipe(recipe) {
        for (let index = 0; index < $storeFE.length; index++) {
            if ($storeFE[index] == recipe) {
                var l = $undo.length;
                $undo[l] = {
                    "task": "Replace",
                    "index": index,
                    "recipe": recipe
                }
                let response = await fetch("https://r7qi88.deta.dev/random");
		        let recipes = await response.json();
                $storeFE[index] = recipes[0];
            } 
        }
        console.log($undo)
    }

    function toggle() {
        visible = !visible;
    }

    export let recipes;
    let visible = false;

	import {
		Container,
		Row,
		Col,
        Icon
	} from 'sveltestrap';

    import { storeFE, undo } from '../scripts/store.js'
    import { slide } from 'svelte/transition';
</script>

<div>
    <div>
        {#await recipes}
            <p>Loading...</p>
        {:then recipes}
            <div class="header">
                <h3>{recipes.title}</h3>
                <p class="time">{recipes.time}</p>
                <button class="button button_normal" on:click={toggle}><Icon name={visible ? "chevron-up" : "chevron-down"} /></button>
                <button class="button button_normal button_right" on:click={removeRecipe(recipes)}><Icon name="trash" /></button>
                <button class="button button_normal button_right" on:click={replaceRecipe(recipes)}><Icon name="arrow-repeat" /></button>
            </div>
            {#if visible}
            <div transition:slide>
                <div class="card_body">
                    <h4>Ingredients</h4>
                    <Container>
                        <Row cols={3}>
                            {#each recipes.ingredients as ingredient}
                                <Col>
                                    <div class="checkbox">
                                        <input id="ingredient" type ="checkbox" >
                                        <label for="ingredient"> {ingredient} </label>
                                    </div>
                                </Col>
                            {/each}
                        </Row>
                    </Container>
                    <h4>Steps</h4>
                    <div>
                        <ol>
                            {#each recipes.steps as step}
                                <div class="step">
                                    <li>{step}</li>
                                </div>
                            {/each}
                        </ol>
                    </div>
                </div>
            </div>
            {/if}
        {/await}
        </div>
</div>

<style>
    .header {
        background: #05386B;
        border-top-left-radius: 8px;
        border-top-right-radius: 8px;
        padding: 10px;
    }
    
    h3 {
        color: #EDF5E1;
        font-size: 16pt;
        margin: 5px;
        margin-bottom: 0;
        padding: 5px;
    }

    .time {
        color: #EDF5E1;
        margin: 5px;
        margin-top: 0;
        padding: 5px;
    }

    .button_normal {
        background: #379683; 
        color: #EDF5E1;
        margin: 10px;
    }

    .checkbox input,
    .checkbox label {
        color: #05386B;
        /* width: auto; */
        margin-left: 0;
        display: inline-block;
        vertical-align: middle;
        /* float: left; */
    }

    ol {
        background: #EDF5E1;
        color: #05386B;
        border-radius: 4px;
        list-style-position: outside;
    }

    /* li {
        margin: 5px;
        margin-top: 10px;
        padding-top: 10px;
    } */

    .card_body {
        background: #8EE4AF;
        padding: 20px;
        border-bottom-left-radius: 8px;
        border-bottom-right-radius: 8px;
    }

    .step {
        padding: 15px;
        width: 98%;
    }

    .step + .step {
        border-top: 1px solid darkgrey;
        padding-top: 10px;
    }

    h4 {
        color: #05386B;
        font-size: large;
    }
</style>
