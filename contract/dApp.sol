pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract DApp1155 is ERC1155 {
	event TransferNFT(address _operator, address _from, address _to, uint _token, uint _amt, string _ext);
	mapping(uint => string) private _tokenURIs;

	constructor() ERC1155("https://baseUrl.com"){}

	function mint(address  _to, uint  _token, uint  _amt, string memory _tokenUrl) public returns(uint _ntoken) {
		_mint(_to, _token,  _amt, "");
		_tokenURIs[_token] = _tokenUrl;
		return _token;
	}

	function transfer(address _from, address _to, uint _token, uint _amt, string memory _ext) public{
		bytes memory data;
		data = bytes(_ext);
		_safeTransferFrom(_from, _to, _token, _amt, data);
		emit TransferNFT(_from, _from, _to, _token, _amt, _ext);
	}
}