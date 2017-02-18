// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestPoolABI is the input ABI used to generate the binding from.
const TestPoolABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"newVersionReleased\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256[]\"}],\"name\":\"array\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"computeLeaf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"shareIndex\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"augCountersBranch\",\"type\":\"uint256[]\"},{\"name\":\"augHashesBranch\",\"type\":\"uint256[]\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"indexInElementsArray\",\"type\":\"uint256\"},{\"name\":\"elements\",\"type\":\"uint256[]\"},{\"name\":\"witness\",\"type\":\"uint256[]\"},{\"name\":\"branchSize\",\"type\":\"uint256\"}],\"name\":\"computeCacheRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"x3\",\"type\":\"uint256\"},{\"name\":\"x4\",\"type\":\"uint256\"}],\"name\":\"hash4\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getClaimSeed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"v1\",\"type\":\"uint256\"},{\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"fnv\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes32\"},{\"name\":\"nonceLe\",\"type\":\"bytes8\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"branchSize\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"hashimoto\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"rlpHeader\",\"type\":\"bytes\"}],\"name\":\"parseBlockHeader_debug\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"merkleRoot\",\"type\":\"uint128\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint64\"},{\"name\":\"branchDepth\",\"type\":\"uint64\"},{\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes\"}],\"name\":\"bytesstream\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"uint256\"},{\"name\":\"nonceLe\",\"type\":\"uint256\"}],\"name\":\"computeS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[16]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"shanumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"input\",\"type\":\"bytes32\"}],\"name\":\"halfsha3\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"declareNewerVersion\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"toBE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"numShares\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\"},{\"name\":\"augMerkle\",\"type\":\"uint256\"}],\"name\":\"submitClaim\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"uint256[16]\"},{\"name\":\"cmix\",\"type\":\"uint256[8]\"}],\"name\":\"computeSha3\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"ErrorLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"msg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"VerifyAgt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"Log\",\"type\":\"event\"}]"

// TestPool is an auto generated Go binding around an Ethereum contract.
type TestPool struct {
	TestPoolCaller     // Read-only binding to the contract
	TestPoolTransactor // Write-only binding to the contract
}

// TestPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPoolSession struct {
	Contract     *TestPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPoolCallerSession struct {
	Contract *TestPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TestPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPoolTransactorSession struct {
	Contract     *TestPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TestPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPoolRaw struct {
	Contract *TestPool // Generic contract binding to access the raw methods on
}

// TestPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPoolCallerRaw struct {
	Contract *TestPoolCaller // Generic read-only contract binding to access the raw methods on
}

// TestPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPoolTransactorRaw struct {
	Contract *TestPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPool creates a new instance of TestPool, bound to a specific deployed contract.
func NewTestPool(address common.Address, backend bind.ContractBackend) (*TestPool, error) {
	contract, err := bindTestPool(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPool{TestPoolCaller: TestPoolCaller{contract: contract}, TestPoolTransactor: TestPoolTransactor{contract: contract}}, nil
}

// NewTestPoolCaller creates a new read-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolCaller(address common.Address, caller bind.ContractCaller) (*TestPoolCaller, error) {
	contract, err := bindTestPool(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TestPoolCaller{contract: contract}, nil
}

// NewTestPoolTransactor creates a new write-only instance of TestPool, bound to a specific deployed contract.
func NewTestPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPoolTransactor, error) {
	contract, err := bindTestPool(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TestPoolTransactor{contract: contract}, nil
}

// bindTestPool binds a generic wrapper to an already deployed contract.
func bindTestPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.TestPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.TestPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPool *TestPoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPool *TestPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPool *TestPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPool.Contract.contract.Transact(opts, method, params...)
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) ComputeCacheRoot(opts *bind.CallOpts, index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "computeCacheRoot", index, indexInElementsArray, elements, witness, branchSize)
	return *ret0, err
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) ComputeCacheRoot(index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeCacheRoot(&_TestPool.CallOpts, index, indexInElementsArray, elements, witness, branchSize)
}

// ComputeCacheRoot is a free data retrieval call binding the contract method 0x3aa4868a.
//
// Solidity: function computeCacheRoot(index uint256, indexInElementsArray uint256, elements uint256[], witness uint256[], branchSize uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) ComputeCacheRoot(index *big.Int, indexInElementsArray *big.Int, elements []*big.Int, witness []*big.Int, branchSize *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeCacheRoot(&_TestPool.CallOpts, index, indexInElementsArray, elements, witness, branchSize)
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) ComputeLeaf(opts *bind.CallOpts, dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "computeLeaf", dataSetLookup, index)
	return *ret0, err
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) ComputeLeaf(dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeLeaf(&_TestPool.CallOpts, dataSetLookup, index)
}

// ComputeLeaf is a free data retrieval call binding the contract method 0x275ccb13.
//
// Solidity: function computeLeaf(dataSetLookup uint256[], index uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) ComputeLeaf(dataSetLookup []*big.Int, index *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeLeaf(&_TestPool.CallOpts, dataSetLookup, index)
}

// ComputeS is a free data retrieval call binding the contract method 0x85e1684c.
//
// Solidity: function computeS(header uint256, nonceLe uint256) constant returns(uint256[16])
func (_TestPool *TestPoolCaller) ComputeS(opts *bind.CallOpts, header *big.Int, nonceLe *big.Int) ([16]*big.Int, error) {
	var (
		ret0 = new([16]*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "computeS", header, nonceLe)
	return *ret0, err
}

// ComputeS is a free data retrieval call binding the contract method 0x85e1684c.
//
// Solidity: function computeS(header uint256, nonceLe uint256) constant returns(uint256[16])
func (_TestPool *TestPoolSession) ComputeS(header *big.Int, nonceLe *big.Int) ([16]*big.Int, error) {
	return _TestPool.Contract.ComputeS(&_TestPool.CallOpts, header, nonceLe)
}

// ComputeS is a free data retrieval call binding the contract method 0x85e1684c.
//
// Solidity: function computeS(header uint256, nonceLe uint256) constant returns(uint256[16])
func (_TestPool *TestPoolCallerSession) ComputeS(header *big.Int, nonceLe *big.Int) ([16]*big.Int, error) {
	return _TestPool.Contract.ComputeS(&_TestPool.CallOpts, header, nonceLe)
}

// ComputeSha3 is a free data retrieval call binding the contract method 0xea135eeb.
//
// Solidity: function computeSha3(s uint256[16], cmix uint256[8]) constant returns(uint256)
func (_TestPool *TestPoolCaller) ComputeSha3(opts *bind.CallOpts, s [16]*big.Int, cmix [8]*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "computeSha3", s, cmix)
	return *ret0, err
}

// ComputeSha3 is a free data retrieval call binding the contract method 0xea135eeb.
//
// Solidity: function computeSha3(s uint256[16], cmix uint256[8]) constant returns(uint256)
func (_TestPool *TestPoolSession) ComputeSha3(s [16]*big.Int, cmix [8]*big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeSha3(&_TestPool.CallOpts, s, cmix)
}

// ComputeSha3 is a free data retrieval call binding the contract method 0xea135eeb.
//
// Solidity: function computeSha3(s uint256[16], cmix uint256[8]) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) ComputeSha3(s [16]*big.Int, cmix [8]*big.Int) (*big.Int, error) {
	return _TestPool.Contract.ComputeSha3(&_TestPool.CallOpts, s, cmix)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() constant returns(address)
func (_TestPool *TestPoolCaller) ContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "contractAddress")
	return *ret0, err
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() constant returns(address)
func (_TestPool *TestPoolSession) ContractAddress() (common.Address, error) {
	return _TestPool.Contract.ContractAddress(&_TestPool.CallOpts)
}

// ContractAddress is a free data retrieval call binding the contract method 0xf6b4dfb4.
//
// Solidity: function contractAddress() constant returns(address)
func (_TestPool *TestPoolCallerSession) ContractAddress() (common.Address, error) {
	return _TestPool.Contract.ContractAddress(&_TestPool.CallOpts)
}

// Fnv is a free data retrieval call binding the contract method 0x48f8fe69.
//
// Solidity: function fnv(v1 uint256, v2 uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Fnv(opts *bind.CallOpts, v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "fnv", v1, v2)
	return *ret0, err
}

// Fnv is a free data retrieval call binding the contract method 0x48f8fe69.
//
// Solidity: function fnv(v1 uint256, v2 uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Fnv(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Fnv(&_TestPool.CallOpts, v1, v2)
}

// Fnv is a free data retrieval call binding the contract method 0x48f8fe69.
//
// Solidity: function fnv(v1 uint256, v2 uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Fnv(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Fnv(&_TestPool.CallOpts, v1, v2)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolCaller) GetClaimSeed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "getClaimSeed")
	return *ret0, err
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolSession) GetClaimSeed() (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts)
}

// GetClaimSeed is a free data retrieval call binding the contract method 0x47740bdc.
//
// Solidity: function getClaimSeed() constant returns(uint256)
func (_TestPool *TestPoolCallerSession) GetClaimSeed() (*big.Int, error) {
	return _TestPool.Contract.GetClaimSeed(&_TestPool.CallOpts)
}

// Halfsha3 is a free data retrieval call binding the contract method 0xa748cf91.
//
// Solidity: function halfsha3(input bytes32) constant returns(uint256[2])
func (_TestPool *TestPoolCaller) Halfsha3(opts *bind.CallOpts, input [32]byte) ([2]*big.Int, error) {
	var (
		ret0 = new([2]*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "halfsha3", input)
	return *ret0, err
}

// Halfsha3 is a free data retrieval call binding the contract method 0xa748cf91.
//
// Solidity: function halfsha3(input bytes32) constant returns(uint256[2])
func (_TestPool *TestPoolSession) Halfsha3(input [32]byte) ([2]*big.Int, error) {
	return _TestPool.Contract.Halfsha3(&_TestPool.CallOpts, input)
}

// Halfsha3 is a free data retrieval call binding the contract method 0xa748cf91.
//
// Solidity: function halfsha3(input bytes32) constant returns(uint256[2])
func (_TestPool *TestPoolCallerSession) Halfsha3(input [32]byte) ([2]*big.Int, error) {
	return _TestPool.Contract.Halfsha3(&_TestPool.CallOpts, input)
}

// Hash4 is a free data retrieval call binding the contract method 0x3b3ccded.
//
// Solidity: function hash4(x1 uint256, x2 uint256, x3 uint256, x4 uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Hash4(opts *bind.CallOpts, x1 *big.Int, x2 *big.Int, x3 *big.Int, x4 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "hash4", x1, x2, x3, x4)
	return *ret0, err
}

// Hash4 is a free data retrieval call binding the contract method 0x3b3ccded.
//
// Solidity: function hash4(x1 uint256, x2 uint256, x3 uint256, x4 uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Hash4(x1 *big.Int, x2 *big.Int, x3 *big.Int, x4 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hash4(&_TestPool.CallOpts, x1, x2, x3, x4)
}

// Hash4 is a free data retrieval call binding the contract method 0x3b3ccded.
//
// Solidity: function hash4(x1 uint256, x2 uint256, x3 uint256, x4 uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Hash4(x1 *big.Int, x2 *big.Int, x3 *big.Int, x4 *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hash4(&_TestPool.CallOpts, x1, x2, x3, x4)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Hashimoto(opts *bind.CallOpts, header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "hashimoto", header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
	return *ret0, err
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// Hashimoto is a free data retrieval call binding the contract method 0x58e69c5a.
//
// Solidity: function hashimoto(header bytes32, nonceLe bytes8, fullSizeIn128Resultion uint256, dataSetLookup uint256[], witnessForLookup uint256[], branchSize uint256, root uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Hashimoto(header [32]byte, nonceLe [8]byte, fullSizeIn128Resultion *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, branchSize *big.Int, root *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Hashimoto(&_TestPool.CallOpts, header, nonceLe, fullSizeIn128Resultion, dataSetLookup, witnessForLookup, branchSize, root)
}

// MinerAddress is a free data retrieval call binding the contract method 0x59275c84.
//
// Solidity: function minerAddress() constant returns(address)
func (_TestPool *TestPoolCaller) MinerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "minerAddress")
	return *ret0, err
}

// MinerAddress is a free data retrieval call binding the contract method 0x59275c84.
//
// Solidity: function minerAddress() constant returns(address)
func (_TestPool *TestPoolSession) MinerAddress() (common.Address, error) {
	return _TestPool.Contract.MinerAddress(&_TestPool.CallOpts)
}

// MinerAddress is a free data retrieval call binding the contract method 0x59275c84.
//
// Solidity: function minerAddress() constant returns(address)
func (_TestPool *TestPoolCallerSession) MinerAddress() (common.Address, error) {
	return _TestPool.Contract.MinerAddress(&_TestPool.CallOpts)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCaller) NewVersionReleased(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "newVersionReleased")
	return *ret0, err
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// NewVersionReleased is a free data retrieval call binding the contract method 0x0289e966.
//
// Solidity: function newVersionReleased() constant returns(bool)
func (_TestPool *TestPoolCallerSession) NewVersionReleased() (bool, error) {
	return _TestPool.Contract.NewVersionReleased(&_TestPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolSession) Owner() (common.Address, error) {
	return _TestPool.Contract.Owner(&_TestPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TestPool *TestPoolCallerSession) Owner() (common.Address, error) {
	return _TestPool.Contract.Owner(&_TestPool.CallOpts)
}

// ParseBlockHeader_debug is a free data retrieval call binding the contract method 0x797c9093.
//
// Solidity: function parseBlockHeader_debug(rlpHeader bytes) constant returns(uint256[5])
func (_TestPool *TestPoolCaller) ParseBlockHeader_debug(opts *bind.CallOpts, rlpHeader []byte) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "parseBlockHeader_debug", rlpHeader)
	return *ret0, err
}

// ParseBlockHeader_debug is a free data retrieval call binding the contract method 0x797c9093.
//
// Solidity: function parseBlockHeader_debug(rlpHeader bytes) constant returns(uint256[5])
func (_TestPool *TestPoolSession) ParseBlockHeader_debug(rlpHeader []byte) ([5]*big.Int, error) {
	return _TestPool.Contract.ParseBlockHeader_debug(&_TestPool.CallOpts, rlpHeader)
}

// ParseBlockHeader_debug is a free data retrieval call binding the contract method 0x797c9093.
//
// Solidity: function parseBlockHeader_debug(rlpHeader bytes) constant returns(uint256[5])
func (_TestPool *TestPoolCallerSession) ParseBlockHeader_debug(rlpHeader []byte) ([5]*big.Int, error) {
	return _TestPool.Contract.ParseBlockHeader_debug(&_TestPool.CallOpts, rlpHeader)
}

// Shanumbers is a free data retrieval call binding the contract method 0x9c73188d.
//
// Solidity: function shanumbers(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) Shanumbers(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "shanumbers", x, y)
	return *ret0, err
}

// Shanumbers is a free data retrieval call binding the contract method 0x9c73188d.
//
// Solidity: function shanumbers(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) Shanumbers(x *big.Int, y *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Shanumbers(&_TestPool.CallOpts, x, y)
}

// Shanumbers is a free data retrieval call binding the contract method 0x9c73188d.
//
// Solidity: function shanumbers(x uint256, y uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) Shanumbers(x *big.Int, y *big.Int) (*big.Int, error) {
	return _TestPool.Contract.Shanumbers(&_TestPool.CallOpts, x, y)
}

// ToBE is a free data retrieval call binding the contract method 0xe422f699.
//
// Solidity: function toBE(x uint256) constant returns(uint256)
func (_TestPool *TestPoolCaller) ToBE(opts *bind.CallOpts, x *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "toBE", x)
	return *ret0, err
}

// ToBE is a free data retrieval call binding the contract method 0xe422f699.
//
// Solidity: function toBE(x uint256) constant returns(uint256)
func (_TestPool *TestPoolSession) ToBE(x *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ToBE(&_TestPool.CallOpts, x)
}

// ToBE is a free data retrieval call binding the contract method 0xe422f699.
//
// Solidity: function toBE(x uint256) constant returns(uint256)
func (_TestPool *TestPoolCallerSession) ToBE(x *big.Int) (*big.Int, error) {
	return _TestPool.Contract.ToBE(&_TestPool.CallOpts, x)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TestPool.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_TestPool *TestPoolCallerSession) Version() (string, error) {
	return _TestPool.Contract.Version(&_TestPool.CallOpts)
}

// Array is a paid mutator transaction binding the contract method 0x1335b639.
//
// Solidity: function array(x uint256[]) returns(bytes32)
func (_TestPool *TestPoolTransactor) Array(opts *bind.TransactOpts, x []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "array", x)
}

// Array is a paid mutator transaction binding the contract method 0x1335b639.
//
// Solidity: function array(x uint256[]) returns(bytes32)
func (_TestPool *TestPoolSession) Array(x []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.Array(&_TestPool.TransactOpts, x)
}

// Array is a paid mutator transaction binding the contract method 0x1335b639.
//
// Solidity: function array(x uint256[]) returns(bytes32)
func (_TestPool *TestPoolTransactorSession) Array(x []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.Array(&_TestPool.TransactOpts, x)
}

// Bytesstream is a paid mutator transaction binding the contract method 0x83006e57.
//
// Solidity: function bytesstream(x bytes) returns(bytes32)
func (_TestPool *TestPoolTransactor) Bytesstream(opts *bind.TransactOpts, x []byte) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "bytesstream", x)
}

// Bytesstream is a paid mutator transaction binding the contract method 0x83006e57.
//
// Solidity: function bytesstream(x bytes) returns(bytes32)
func (_TestPool *TestPoolSession) Bytesstream(x []byte) (*types.Transaction, error) {
	return _TestPool.Contract.Bytesstream(&_TestPool.TransactOpts, x)
}

// Bytesstream is a paid mutator transaction binding the contract method 0x83006e57.
//
// Solidity: function bytesstream(x bytes) returns(bytes32)
func (_TestPool *TestPoolTransactorSession) Bytesstream(x []byte) (*types.Transaction, error) {
	return _TestPool.Contract.Bytesstream(&_TestPool.TransactOpts, x)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactor) DeclareNewerVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "declareNewerVersion")
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// DeclareNewerVersion is a paid mutator transaction binding the contract method 0xe3d86998.
//
// Solidity: function declareNewerVersion() returns()
func (_TestPool *TestPoolTransactorSession) DeclareNewerVersion() (*types.Transaction, error) {
	return _TestPool.Contract.DeclareNewerVersion(&_TestPool.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(timestamp uint256) returns()
func (_TestPool *TestPoolTransactor) Register(opts *bind.TransactOpts, timestamp *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "register", timestamp)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(timestamp uint256) returns()
func (_TestPool *TestPoolSession) Register(timestamp *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, timestamp)
}

// Register is a paid mutator transaction binding the contract method 0xf207564e.
//
// Solidity: function register(timestamp uint256) returns()
func (_TestPool *TestPoolTransactorSession) Register(timestamp *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.Register(&_TestPool.TransactOpts, timestamp)
}

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolTransactor) SetEpochData(opts *bind.TransactOpts, merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setEpochData", merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolSession) SetEpochData(merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetEpochData is a paid mutator transaction binding the contract method 0x820a22c0.
//
// Solidity: function setEpochData(merkleRoot uint128, fullSizeIn128Resultion uint64, branchDepth uint64, epoch uint256) returns()
func (_TestPool *TestPoolTransactorSession) SetEpochData(merkleRoot *big.Int, fullSizeIn128Resultion uint64, branchDepth uint64, epoch *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetEpochData(&_TestPool.TransactOpts, merkleRoot, fullSizeIn128Resultion, branchDepth, epoch)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x5d5e8875.
//
// Solidity: function setMerkleRoot(root uint256) returns()
func (_TestPool *TestPoolTransactor) SetMerkleRoot(opts *bind.TransactOpts, root *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "setMerkleRoot", root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x5d5e8875.
//
// Solidity: function setMerkleRoot(root uint256) returns()
func (_TestPool *TestPoolSession) SetMerkleRoot(root *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetMerkleRoot(&_TestPool.TransactOpts, root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x5d5e8875.
//
// Solidity: function setMerkleRoot(root uint256) returns()
func (_TestPool *TestPoolTransactorSession) SetMerkleRoot(root *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SetMerkleRoot(&_TestPool.TransactOpts, root)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactor) SubmitClaim(opts *bind.TransactOpts, numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "submitClaim", numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// SubmitClaim is a paid mutator transaction binding the contract method 0xe7dac983.
//
// Solidity: function submitClaim(numShares uint256, difficulty uint256, min uint256, max uint256, augMerkle uint256) returns()
func (_TestPool *TestPoolTransactorSession) SubmitClaim(numShares *big.Int, difficulty *big.Int, min *big.Int, max *big.Int, augMerkle *big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.SubmitClaim(&_TestPool.TransactOpts, numShares, difficulty, min, max, augMerkle)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactor) VerifyClaim(opts *bind.TransactOpts, rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.contract.Transact(opts, "verifyClaim", rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}

// VerifyClaim is a paid mutator transaction binding the contract method 0x35ffbe74.
//
// Solidity: function verifyClaim(rlpHeader bytes, nonce uint256, shareIndex uint256, dataSetLookup uint256[], witnessForLookup uint256[], augCountersBranch uint256[], augHashesBranch uint256[]) returns(uint256)
func (_TestPool *TestPoolTransactorSession) VerifyClaim(rlpHeader []byte, nonce *big.Int, shareIndex *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int, augCountersBranch []*big.Int, augHashesBranch []*big.Int) (*types.Transaction, error) {
	return _TestPool.Contract.VerifyClaim(&_TestPool.TransactOpts, rlpHeader, nonce, shareIndex, dataSetLookup, witnessForLookup, augCountersBranch, augHashesBranch)
}
