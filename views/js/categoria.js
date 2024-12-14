const bntCategoria = document.getElementById("bnt-categoria")
const formContainer = document.getElementById("container-produtos")

const bntProduto = document.getElementById("bnt-produto")


function categoryForm(){

    formContainer.innerHTML = `

    <h3>Formulário de Categoria</h3>

    <form>
        <label for="category-name"  style="margin-left: 2rem;">Nome:</label>
        <input type="text" id="category-name" name="category-name" required>
        <button type="submit" class="btn btn-success">Salvar</button>
      </form>
    
    `;
}

function produtoForm(){

    formContainer.innerHTML = `

            <h3>Formulário de Produto</h3>
            <form class="row g-3" style="margin-top: 2rem;">

                <div class="col-md-4">
                    <label for="product-name" style="margin-left: 1rem; margin-right: 1rem;">Nome:</label>
                    <input type="text" id="product-name" name="product-name" required>
                </div>

                <div class="col-md-4">
                 <label for="product-price"  style="margin-left: 1rem; margin-right: 2rem;">Preço:</label>
                 <input type="number" id="product-price" name="product-price" required>
                </div>

                <div class="col-md-4">
                  <label for="product-lot"  style="margin-left: 1rem; margin-right: 1rem;">Lote:</label>
                  <input type="text" id="product-lot" name="product-lot" required>
                </div>

                <div class="col-md-4">
                  <label for="product-lot"   style="margin-left:1rem;">Validade:</label>
                  <input type="text" id="product-lot" name="product-lot" required>
                </div>

                <div class="col-md-4">
                  <label for="product-lot" style="margin-left: 1rem;">Categoria:</label>
                  <input type="text" id="product-lot" name="product-lot" required>
                </div>


                <div>
                  <button type="submit" class="btn btn-success" style="margin-left: 25rem; margin-top: 3rem;">Salvar</button>
                </div>
             
            </form>
    
    `;
}

bntProduto.addEventListener("click", produtoForm)
bntCategoria.addEventListener("click", categoryForm)