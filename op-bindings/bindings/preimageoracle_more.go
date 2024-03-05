// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const PreimageOracleStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageLengths\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimageParts\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1002,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"preimagePartOk\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\"},{\"astId\":1003,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"zeroHashes\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_bytes32)16_storage\"},{\"astId\":1004,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"19\",\"type\":\"t_array(t_struct(LargePreimageProposalKeys)1010_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBranches\",\"offset\":0,\"slot\":\"20\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\"},{\"astId\":1006,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalMetadata\",\"offset\":0,\"slot\":\"21\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1011))\"},{\"astId\":1007,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBonds\",\"offset\":0,\"slot\":\"22\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_uint256))\"},{\"astId\":1008,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalParts\",\"offset\":0,\"slot\":\"23\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\"},{\"astId\":1009,\"contract\":\"src/cannon/PreimageOracle.sol:PreimageOracle\",\"label\":\"proposalBlocks\",\"offset\":0,\"slot\":\"24\",\"type\":\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_bytes32)16_storage\":{\"encoding\":\"inplace\",\"label\":\"bytes32[16]\",\"numberOfBytes\":\"512\",\"base\":\"t_bytes32\"},\"t_array(t_struct(LargePreimageProposalKeys)1010_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(LargePreimageProposalKeys)1010_storage\"},\"t_array(t_uint64)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint64[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint64\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_bytes32)16_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32[16]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_array(t_uint64)dyn_storage))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e uint64[]))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_uint256)\"},\"t_mapping(t_address,t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1011))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(uint256 =\u003e LPPMetaData))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1011)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bool)\"},\"t_mapping(t_bytes32,t_mapping(t_uint256,t_bytes32))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(uint256 =\u003e bytes32))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_uint256,t_bytes32)\"},\"t_mapping(t_bytes32,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_array(t_bytes32)16_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32[16])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_bytes32)16_storage\"},\"t_mapping(t_uint256,t_array(t_uint64)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint64[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint64)dyn_storage\"},\"t_mapping(t_uint256,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bool\"},\"t_mapping(t_uint256,t_bytes32)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e bytes32)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_bytes32\"},\"t_mapping(t_uint256,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_uint256\"},\"t_mapping(t_uint256,t_userDefinedValueType(LPPMetaData)1011)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e LPPMetaData)\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_userDefinedValueType(LPPMetaData)1011\"},\"t_struct(LargePreimageProposalKeys)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"struct PreimageOracle.LargePreimageProposalKeys\",\"numberOfBytes\":\"64\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_userDefinedValueType(LPPMetaData)1011\":{\"encoding\":\"inplace\",\"label\":\"LPPMetaData\",\"numberOfBytes\":\"32\"}}}"

var PreimageOracleStorageLayout = new(solc.StorageLayout)

var PreimageOracleDeployedBin = "0x6080604052600436106101cd5760003560e01c80638dc4be11116100f7578063dd24f9bf11610095578063ec5efcbc11610064578063ec5efcbc1461065f578063f3f480d91461067f578063faf37bc7146106b2578063fef2b4ed146106c557600080fd5b8063dd24f9bf1461059f578063ddcd58de146105d2578063e03110e11461060a578063e15926111461063f57600080fd5b8063b2e67ba8116100d1578063b2e67ba814610512578063b4801e611461054a578063d18534b51461056a578063da35c6641461058a57600080fd5b80638dc4be11146104835780639d53a648146104a35780639d7e8769146104f257600080fd5b806354fd4d501161016f5780637917de1d1161013e5780637917de1d146103bf5780637ac54767146103df5780638542cf50146103ff578063882856ef1461044a57600080fd5b806354fd4d50146102dd57806361238bde146103335780636551927b1461036b5780637051472e146103a357600080fd5b80632055b36b116101ab5780632055b36b146102735780633909af5c146102885780634d52b4c9146102a857806352f0f3ad146102bd57600080fd5b8063013cf08b146101d25780630359a5631461022357806304697c7814610251575b600080fd5b3480156101de57600080fd5b506101f26101ed366004612dd0565b6106f2565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152015b60405180910390f35b34801561022f57600080fd5b5061024361023e366004612e12565b610737565b60405190815260200161021a565b34801561025d57600080fd5b5061027161026c366004612e85565b61086f565b005b34801561027f57600080fd5b50610243601081565b34801561029457600080fd5b506102716102a33660046130a9565b6109ff565b3480156102b457600080fd5b50610243610c56565b3480156102c957600080fd5b506102436102d8366004613195565b610c71565b3480156102e957600080fd5b506103266040518060400160405280600581526020017f302e312e3000000000000000000000000000000000000000000000000000000081525081565b60405161021a91906131fc565b34801561033f57600080fd5b5061024361034e36600461324d565b600160209081526000928352604080842090915290825290205481565b34801561037757600080fd5b50610243610386366004612e12565b601560209081526000928352604080842090915290825290205481565b3480156103af57600080fd5b506102436703782dace9d9000081565b3480156103cb57600080fd5b506102716103da36600461326f565b610d46565b3480156103eb57600080fd5b506102436103fa366004612dd0565b611236565b34801561040b57600080fd5b5061043a61041a36600461324d565b600260209081526000928352604080842090915290825290205460ff1681565b604051901515815260200161021a565b34801561045657600080fd5b5061046a61046536600461330b565b61124d565b60405167ffffffffffffffff909116815260200161021a565b34801561048f57600080fd5b5061027161049e36600461333e565b6112a7565b3480156104af57600080fd5b506102436104be366004612e12565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152601860209081526040808320938352929052205490565b3480156104fe57600080fd5b5061027161050d36600461338a565b6113a2565b34801561051e57600080fd5b5061024361052d366004612e12565b601760209081526000928352604080842090915290825290205481565b34801561055657600080fd5b5061024361056536600461330b565b6115b3565b34801561057657600080fd5b506102716105853660046130a9565b6115e5565b34801561059657600080fd5b50601354610243565b3480156105ab57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610243565b3480156105de57600080fd5b506102436105ed366004612e12565b601660209081526000928352604080842090915290825290205481565b34801561061657600080fd5b5061062a61062536600461324d565b6119a7565b6040805192835260208301919091520161021a565b34801561064b57600080fd5b5061027161065a36600461333e565b611a98565b34801561066b57600080fd5b5061027161067a366004613416565b611ba0565b34801561068b57600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610243565b6102716106c03660046134af565b611d26565b3480156106d157600080fd5b506102436106e0366004612dd0565b60006020819052908152604090205481565b6013818154811061070257600080fd5b60009182526020909120600290910201805460019091015473ffffffffffffffffffffffffffffffffffffffff909116915082565b73ffffffffffffffffffffffffffffffffffffffff82166000908152601560209081526040808320848452909152812054819061077a9060601c63ffffffff1690565b63ffffffff16905060005b6010811015610867578160011660010361080d5773ffffffffffffffffffffffffffffffffffffffff85166000908152601460209081526040808320878452909152902081601081106107da576107da6134eb565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925061084e565b8260038260108110610821576108216134eb565b01546040805160208101939093528201526060016040516020818303038152906040528051906020012092505b60019190911c908061085f81613549565b915050610785565b505092915050565b7f00000000000000000000000000000000000000000000000000000000000000004210156108c9576040517f299f254900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080600080608060146030823785878260140137601480870182207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f06000000000000000000000000000000000000000000000000000000000000001794506000908190889084018b5afa94503d60010191506008820189106109565763fe2549876000526004601cfd5b60c082901b81526008018481533d6000600183013e88017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8015160008481526002602090815260408083208c8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915587845282528083209b83529a81528a82209290925593845283905296909120959095555050505050565b6000610a0b8a8a610737565b9050610a2e86868360208b0135610a29610a248d613581565b611f91565b611fd1565b8015610a4c5750610a4c8383836020880135610a29610a248a613581565b610a82576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b866040013588604051602001610a989190613650565b6040516020818303038152906040528051906020012014610ae5576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836020013587602001356001610afb919061368e565b14610b32576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b7a88610b4086806136a6565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061203292505050565b610b838861218d565b836040013588604051602001610b999190613650565b6040516020818303038152906040528051906020012003610be6576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8a1660009081526015602090815260408083208c8452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001179055610c4a8a8a33612935565b50505050505050505050565b6001610c646010600261382d565b610c6e9190613839565b81565b6000610c7d86866129ee565b9050610c8a83600861368e565b821180610c975750602083115b15610cce576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000602081815260c085901b82526008959095528251828252600286526040808320858452875280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558484528752808320948352938652838220558181529384905292205592915050565b60608115610d5f57610d588686612a9b565b9050610d99565b85858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509293505050505b3360009081526014602090815260408083208b845290915280822081516102008101928390529160109082845b815481526020019060010190808311610dc657505050505090506000601560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008b81526020019081526020016000205490506000610e478260601c63ffffffff1690565b63ffffffff169050333214610e88576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e988260801c63ffffffff1690565b63ffffffff16600003610ed7576040517f87138d5c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ee18260c01c90565b67ffffffffffffffff1615610f22576040517f475a253500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b898114610f5b576040517f60f95d5a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f6889898d8886612b14565b83516020850160888204881415608883061715610f8d576307b1daf16000526004601cfd5b60405160c8810160405260005b8381101561103d578083018051835260208101516020840152604081015160408401526060810151606084015260808101516080840152508460888301526088810460051b8b013560a883015260c882206001860195508560005b6102008110156110325760018216156110125782818b0152611032565b8981015160009081526020938452604090209260019290921c9101610ff5565b505050608801610f9a565b50505050600160106002611051919061382d565b61105b9190613839565b811115611094576040517f6229572300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111096110a78360401c63ffffffff1690565b6110b79063ffffffff168a61368e565b60401b7fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff606084901b167fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff8516171790565b915084156111965777ffffffffffffffffffffffffffffffffffffffffffffffff82164260c01b1791506111438260801c63ffffffff1690565b63ffffffff166111598360401c63ffffffff1690565b63ffffffff1614611196576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526014602090815260408083208e845290915290206111bc90846010612d46565b50503360008181526018602090815260408083208e8452825280832080546001810182559084528284206004820401805460039092166008026101000a67ffffffffffffffff818102199093164390931602919091179055928252601581528282209c82529b909b52909920989098555050505050505050565b6003816010811061124657600080fd5b0154905081565b6018602052826000526040600020602052816000526040600020818154811061127557600080fd5b906000526020600020906004918282040191900660080292509250509054906101000a900467ffffffffffffffff1681565b6044356000806008830186106112c55763fe2549876000526004601cfd5b60c083901b60805260888386823786600882030151915060206000858360025afa9050806112f257600080fd5b50600080517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0400000000000000000000000000000000000000000000000000000000000000178082526002602090815260408084208a8552825280842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091558385528252808420998452988152888320939093558152908190529490942055505050565b7f00000000000000000000000000000000000000000000000000000000000000004210156113fc576040517f299f254900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080603087600037602060006030600060025afa806114245763f91129696000526004601cfd5b6000517effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f010000000000000000000000000000000000000000000000000000000000000017608081815260a08c905260c08b905260308a60e037603088609083013760008060c083600a5afa9250826114a6576309bde3396000526004601cfd5b602886106114bc5763fe2549876000526004601cfd5b6000602882015278200000000000000000000000000000000000000000000000008152600881018b905285810151935060308a8237603081019b909b52505060509098207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0500000000000000000000000000000000000000000000000000000000000000176000818152600260209081526040808320868452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209583529481528482209a909a559081528089529190912096909655505050505050565b601460205282600052604060002060205281600052604060002081601081106115db57600080fd5b0154925083915050565b73ffffffffffffffffffffffffffffffffffffffff891660009081526015602090815260408083208b845290915290205467ffffffffffffffff811615611658576040517fc334f06900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006116838260c01c90565b6116979067ffffffffffffffff1642613839565b116116ce576040517f55d4cbf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006116da8b8b610737565b90506116f387878360208c0135610a29610a248e613581565b801561171157506117118484836020890135610a29610a248b613581565b611747576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b87604001358960405160200161175d9190613650565b60405160208183030381529060405280519060200120146117aa576040517f1968a90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8460200135886020013560016117c0919061368e565b1415806117f2575060016117da8360601c63ffffffff1690565b6117e49190613850565b63ffffffff16856020013514155b15611829576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61183789610b4087806136a6565b6118408961218d565b600061184b8a612c67565b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f020000000000000000000000000000000000000000000000000000000000000017905060006118a28460a01c63ffffffff1690565b67ffffffffffffffff169050600160026000848152602001908152602001600020600083815260200190815260200160002060006101000a81548160ff021916908315150217905550601760008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008d815260200190815260200160002054600160008481526020019081526020016000206000838152602001908152602001600020819055506119748460801c63ffffffff1690565b600083815260208190526040902063ffffffff9190911690556119988d8d81612935565b50505050505050505050505050565b6000828152600260209081526040808320848452909152812054819060ff16611a30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f7072652d696d616765206d757374206578697374000000000000000000000000604482015260640160405180910390fd5b5060008381526020818152604090912054611a4c81600861368e565b611a5785602061368e565b10611a755783611a6882600861368e565b611a729190613839565b91505b506000938452600160209081526040808620948652939052919092205492909150565b604435600080600883018610611ab65763fe2549876000526004601cfd5b60c083901b6080526088838682378087017ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80151908490207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f02000000000000000000000000000000000000000000000000000000000000001760008181526002602090815260408083208b8452825280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584845282528083209a83529981528982209390935590815290819052959095209190915550505050565b6000611bac8686610737565b9050611bc58383836020880135610a29610a248a613581565b611bfb576040517f09bde33900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b602084013515611c37576040517f9a3b119900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c3f612d84565b611c4d81610b4087806136a6565b611c568161218d565b846040013581604051602001611c6c9190613650565b6040516020818303038152906040528051906020012003611cb9576040517f9843145b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff87166000908152601560209081526040808320898452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000166001179055611d1d878733612935565b50505050505050565b6703782dace9d90000341015611d68576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b333214611da1576040517fba092d1600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611dac816008613875565b63ffffffff168263ffffffff1610611df0576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000008163ffffffff161015611e50576040517f7b1dafd100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336000818152601560209081526040808320878452825280832080547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff1660a09790971b7fffffffffffffffffffffffff00000000ffffffffffffffffffffffffffffffff169690961760809590951b949094179094558251808401845282815280850186815260138054600181018255908452915160029092027f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0908101805473ffffffffffffffffffffffffffffffffffffffff9094167fffffffffffffffffffffffff000000000000000000000000000000000000000090941693909317909255517f66de8ffda797e3de9c05e8fc57b3bf0ec28a930d40b0d285d93c06501cf6a0919091015590815260168352818120938152929091529020349055565b6000816000015182602001518360400151604051602001611fb49392919061389d565b604051602081830303815290604052805190602001209050919050565b60008160005b6010811015612025578060051b880135600186831c166001811461200a576000848152602083905260409020935061201b565b600082815260208590526040902093505b5050600101611fd7565b5090931495945050505050565b608881511461204057600080fd5b60208101602083016120c1565b8260031b8201518060001a8160011a60081b178160021a60101b8260031a60181b17178160041a60201b8260051a60281b178260061a60301b8360071a60381b17171790506120bb816120a6868560059190911b015190565b1867ffffffffffffffff16600586901b840152565b50505050565b6120cd6000838361204d565b6120d96001838361204d565b6120e56002838361204d565b6120f16003838361204d565b6120fd6004838361204d565b6121096005838361204d565b6121156006838361204d565b6121216007838361204d565b61212d6008838361204d565b6121396009838361204d565b612145600a838361204d565b612151600b838361204d565b61215d600c838361204d565b612169600d838361204d565b612175600e838361204d565b612181600f838361204d565b6120bb6010838361204d565b6040805178010000000000008082800000000000808a8000000080008000602082015279808b00000000800000018000000080008081800000000000800991810191909152788a00000000000000880000000080008009000000008000000a60608201527b8000808b800000000000008b8000000000008089800000000000800360808201527f80000000000080028000000000000080000000000000800a800000008000000a60a08201527f800000008000808180000000000080800000000080000001800000008000800860c082015260009060e00160405160208183030381529060405290506020820160208201612815565b6102808101516101e082015161014083015160a0840151845118189118186102a082015161020083015161016084015160c0850151602086015118189118186102c083015161022084015161018085015160e0860151604087015118189118186102e08401516102408501516101a0860151610100870151606088015118189118186103008501516102608601516101c0870151610120880151608089015118189118188084603f1c6123408660011b67ffffffffffffffff1690565b18188584603f1c61235b8660011b67ffffffffffffffff1690565b18188584603f1c6123768660011b67ffffffffffffffff1690565b181895508483603f1c6123938560011b67ffffffffffffffff1690565b181894508387603f1c6123b08960011b67ffffffffffffffff1690565b60208b01518b51861867ffffffffffffffff168c5291189190911897508118600181901b603f9190911c18935060c08801518118601481901c602c9190911b1867ffffffffffffffff1660208901526101208801518718602c81901c60149190911b1867ffffffffffffffff1660c08901526102c08801518618600381901c603d9190911b1867ffffffffffffffff166101208901526101c08801518718601981901c60279190911b1867ffffffffffffffff166102c08901526102808801518218602e81901c60129190911b1867ffffffffffffffff166101c089015260408801518618600281901c603e9190911b1867ffffffffffffffff166102808901526101808801518618601581901c602b9190911b1867ffffffffffffffff1660408901526101a08801518518602781901c60199190911b1867ffffffffffffffff166101808901526102608801518718603881901c60089190911b1867ffffffffffffffff166101a08901526102e08801518518600881901c60389190911b1867ffffffffffffffff166102608901526101e08801518218601781901c60299190911b1867ffffffffffffffff166102e089015260808801518718602581901c601b9190911b1867ffffffffffffffff166101e08901526103008801518718603281901c600e9190911b1867ffffffffffffffff1660808901526102a08801518118603e81901c60029190911b1867ffffffffffffffff166103008901526101008801518518600981901c60379190911b1867ffffffffffffffff166102a08901526102008801518118601381901c602d9190911b1867ffffffffffffffff1661010089015260a08801518218601c81901c60249190911b1867ffffffffffffffff1661020089015260608801518518602481901c601c9190911b1867ffffffffffffffff1660a08901526102408801518518602b81901c60159190911b1867ffffffffffffffff1660608901526102208801518618603181901c600f9190911b1867ffffffffffffffff166102408901526101608801518118603681901c600a9190911b1867ffffffffffffffff166102208901525060e08701518518603a81901c60069190911b1867ffffffffffffffff166101608801526101408701518118603d81901c60039190911b1867ffffffffffffffff1660e0880152505067ffffffffffffffff81166101408601525b5050505050565b600582811b8201805160018501831b8401805160028701851b8601805160038901871b8801805160048b0190981b8901805167ffffffffffffffff861985168918811690995283198a16861889169096528819861683188816909352841986168818871690528419831684189095169052919391929190611d1d565b6127af600082612728565b6127ba600582612728565b6127c5600a82612728565b6127d0600f82612728565b6127db601482612728565b50565b6127e781612283565b6127f0816127a4565b600383901b820151815160c09190911c906120bb90821867ffffffffffffffff168352565b612821600082846127de565b61282d600182846127de565b612839600282846127de565b612845600382846127de565b612851600482846127de565b61285d600582846127de565b612869600682846127de565b612875600782846127de565b612881600882846127de565b61288d600982846127de565b612899600a82846127de565b6128a5600b82846127de565b6128b1600c82846127de565b6128bd600d82846127de565b6128c9600e82846127de565b6128d5600f82846127de565b6128e1601082846127de565b6128ed601182846127de565b6128f9601282846127de565b612905601382846127de565b612911601482846127de565b61291d601582846127de565b612929601682846127de565b6120bb601782846127de565b73ffffffffffffffffffffffffffffffffffffffff83811660009081526016602090815260408083208684529091528082208054908390559051909284169083908381818185875af1925050503d80600081146129ae576040519150601f19603f3d011682016040523d82523d6000602084013e6129b3565b606091505b5050905080612721576040517f83e6cc6b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f01000000000000000000000000000000000000000000000000000000000000007effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831617612a94818360408051600093845233602052918152606090922091527effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001790565b9392505050565b6060604051905081602082018181018286833760888306808015612ae45760888290038501848101848103803687375060806001820353506001845160001a1784538652612afb565b608836843760018353608060878401536088850186525b5050505050601f19603f82510116810160405292915050565b6000612b268260a01c63ffffffff1690565b67ffffffffffffffff1690506000612b448360801c63ffffffff1690565b63ffffffff1690506000612b5e8460401c63ffffffff1690565b63ffffffff169050600883108015612b74575080155b15612ba85760c082901b6000908152883560085283513382526017602090815260408084208a855290915290912055612c5d565b60088310158015612bc6575080612bc0600885613839565b93508310155b8015612bda5750612bd7878261368e565b83105b15612c5d576000612beb8285613839565b905087612bf982602061368e565b10158015612c05575085155b15612c3c576040517ffe25498700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526017602090815260408083208a845290915290209089013590555b5050505050505050565b6000612cea565b66ff00ff00ff00ff8160081c1667ff00ff00ff00ff00612c988360081b67ffffffffffffffff1690565b1617905065ffff0000ffff8160101c1667ffff0000ffff0000612cc58360101b67ffffffffffffffff1690565b1617905060008160201c612ce38360201b67ffffffffffffffff1690565b1792915050565b60808201516020830190612d0290612c6e565b612c6e565b6040820151612d1090612c6e565b60401b17612d28612cfd60018460059190911b015190565b825160809190911b90612d3a90612c6e565b60c01b17179392505050565b8260108101928215612d74579160200282015b82811115612d74578251825591602001919060010190612d59565b50612d80929150612d9c565b5090565b6040518060200160405280612d97612db1565b905290565b5b80821115612d805760008155600101612d9d565b6040518061032001604052806019906020820280368337509192915050565b600060208284031215612de257600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114612e0d57600080fd5b919050565b60008060408385031215612e2557600080fd5b612e2e83612de9565b946020939093013593505050565b60008083601f840112612e4e57600080fd5b50813567ffffffffffffffff811115612e6657600080fd5b602083019150836020828501011115612e7e57600080fd5b9250929050565b60008060008060608587031215612e9b57600080fd5b84359350612eab60208601612de9565b9250604085013567ffffffffffffffff811115612ec757600080fd5b612ed387828801612e3c565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610320810167ffffffffffffffff81118282101715612f3257612f32612edf565b60405290565b6040516060810167ffffffffffffffff81118282101715612f3257612f32612edf565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715612fa257612fa2612edf565b604052919050565b6000610320808385031215612fbe57600080fd5b604051602080820167ffffffffffffffff8382108183111715612fe357612fe3612edf565b8160405283955087601f880112612ff957600080fd5b613001612f0e565b948701949150818886111561301557600080fd5b875b8681101561303d57803583811681146130305760008081fd5b8452928401928401613017565b50909352509295945050505050565b60006060828403121561305e57600080fd5b50919050565b60008083601f84011261307657600080fd5b50813567ffffffffffffffff81111561308e57600080fd5b6020830191508360208260051b8501011115612e7e57600080fd5b60008060008060008060008060006103e08a8c0312156130c857600080fd5b6130d18a612de9565b985060208a013597506130e78b60408c01612faa565b96506103608a013567ffffffffffffffff8082111561310557600080fd5b6131118d838e0161304c565b97506103808c013591508082111561312857600080fd5b6131348d838e01613064565b90975095506103a08c013591508082111561314e57600080fd5b61315a8d838e0161304c565b94506103c08c013591508082111561317157600080fd5b5061317e8c828d01613064565b915080935050809150509295985092959850929598565b600080600080600060a086880312156131ad57600080fd5b505083359560208501359550604085013594606081013594506080013592509050565b60005b838110156131eb5781810151838201526020016131d3565b838111156120bb5750506000910152565b602081526000825180602084015261321b8160408501602087016131d0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6000806040838503121561326057600080fd5b50508035926020909101359150565b600080600080600080600060a0888a03121561328a57600080fd5b8735965060208801359550604088013567ffffffffffffffff808211156132b057600080fd5b6132bc8b838c01612e3c565b909750955060608a01359150808211156132d557600080fd5b506132e28a828b01613064565b909450925050608088013580151581146132fb57600080fd5b8091505092959891949750929550565b60008060006060848603121561332057600080fd5b61332984612de9565b95602085013595506040909401359392505050565b60008060006040848603121561335357600080fd5b83359250602084013567ffffffffffffffff81111561337157600080fd5b61337d86828701612e3c565b9497909650939450505050565b600080600080600080600060a0888a0312156133a557600080fd5b8735965060208801359550604088013567ffffffffffffffff808211156133cb57600080fd5b6133d78b838c01612e3c565b909750955060608a01359150808211156133f057600080fd5b506133fd8a828b01612e3c565b989b979a50959894979596608090950135949350505050565b60008060008060006080868803121561342e57600080fd5b61343786612de9565b945060208601359350604086013567ffffffffffffffff8082111561345b57600080fd5b61346789838a0161304c565b9450606088013591508082111561347d57600080fd5b5061348a88828901613064565b969995985093965092949392505050565b803563ffffffff81168114612e0d57600080fd5b6000806000606084860312156134c457600080fd5b833592506134d46020850161349b565b91506134e26040850161349b565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361357a5761357a61351a565b5060010190565b60006060823603121561359357600080fd5b61359b612f38565b823567ffffffffffffffff808211156135b357600080fd5b9084019036601f8301126135c657600080fd5b81356020828211156135da576135da612edf565b61360a817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85011601612f5b565b9250818352368183860101111561362057600080fd5b81818501828501376000918301810191909152908352848101359083015250604092830135928101929092525090565b81516103208201908260005b601981101561368557825167ffffffffffffffff1682526020928301929091019060010161365c565b50505092915050565b600082198211156136a1576136a161351a565b500190565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126136db57600080fd5b83018035915067ffffffffffffffff8211156136f657600080fd5b602001915036819003821315612e7e57600080fd5b600181815b8085111561376457817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561374a5761374a61351a565b8085161561375757918102915b93841c9390800290613710565b509250929050565b60008261377b57506001613827565b8161378857506000613827565b816001811461379e57600281146137a8576137c4565b6001915050613827565b60ff8411156137b9576137b961351a565b50506001821b613827565b5060208310610133831016604e8410600b84101617156137e7575081810a613827565b6137f1838361370b565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156138235761382361351a565b0290505b92915050565b6000612a94838361376c565b60008282101561384b5761384b61351a565b500390565b600063ffffffff8381169083168181101561386d5761386d61351a565b039392505050565b600063ffffffff8083168185168083038211156138945761389461351a565b01949350505050565b600084516138af8184602089016131d0565b9190910192835250602082015260400191905056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(PreimageOracleStorageLayoutJSON), PreimageOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["PreimageOracle"] = PreimageOracleStorageLayout
	deployedBytecodes["PreimageOracle"] = PreimageOracleDeployedBin
	immutableReferences["PreimageOracle"] = true
}
