pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract DApp1155 is ERC1155 {

	mapping(uint => string) private _tokenURIs;

	constructor() ERC1155("https://baseUrl.com"){}

	function mint(address _to, uint _token, uint _amt, string _tokenUrl) public returns(uint _token) {
		_mint(_to, _token,  _amt, "");
		_tokenURIs[_token] = _tokenUrl;
		return _token
	}

	function transfer(address _from, address _to, uint _token, uint _amt, string ext) public{
		bytes data
		data = ext
		_safeTransFrom(_from, _to, _token, _amt, data)
	}
}
