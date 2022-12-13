const Product = artifacts.require("Product");

/*
 * uncomment accounts to access the test accounts made available by the
 * Ethereum client
 * See docs: https://www.trufflesuite.com/docs/truffle/testing/writing-tests-in-javascript
 */
contract("Product", async (accounts) => {
  let [fromUser, toUser] = accounts;

  describe("Deployment", function () {
    it("Should deploy with the right greeting", async function () {
      const product = await Product.new();

      const name = "product1";
      const sku = "1";
      const description = "test test";
      const imageURL = "https://test.com";
      const price = 2;
      const stock = 10;
      const result = await product.createProduct(
        name,
        sku,
        description,
        imageURL,
        price,
        stock,
        { from: fromUser }
      );
      assert.equal(result.logs[0].args.name, name);
      assert.equal(result.logs[0].args.description, description);
      assert.equal(result.logs[0].args.imageURL, imageURL);
    });
  });
});
