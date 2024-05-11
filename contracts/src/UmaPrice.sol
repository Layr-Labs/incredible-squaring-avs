// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;


import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/FinderInterface.sol";
import "./interfaces/OptimisticOracleV2Interface.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

// Uma Integration. 

contract PriceFeedAdapterV2 is Ownable {
   
   FinderInterface public finder; //0xf4C48eDAd256326086AEfbd1A53e1896815F8f13
//0x9923D42eF695B5dd9911D05Ac944d4cAca3c4EAB goerli
    //OptimisticOracleV2Interface public oo;
     
     OptimisticOracleV2Interface oo = OptimisticOracleV2Interface(0x9f1263B8f0355673619168b5B8c0248f1d03e88C); // sepolia 

    IERC20 public currency;

    uint256 public umaOraclePrice;

    uint256 public requestTime = 0;
    uint256 public requestTime0 = 0;
    
    

    constructor(
        
        address _finderAddress
       // bytes32 _priceIdentifier
        
    )  {
        finder = FinderInterface(_finderAddress); //(PriceFeed) finder  address sepolia. //0xf4C48eDAd256326086AEfbd1A53e1896815F8f13
        
       // priceIdentifier = _priceIdentifier;
       // https://docs.uma.xyz/resources/approved-price-identifiers
    }


    function _getOptimisticOracle() internal view returns (OptimisticOracleV2Interface) {
        return OptimisticOracleV2Interface(finder.getImplementationAddress("OptimisticOracleV2")); // TODO OracleInterfaces.OptimisticOracleV2
    }

   

     function _getUmaOraclePrice(bytes32 _priceIdentifier) external onlyOwner returns (uint256) {
        OptimisticOracleV2Interface oracle = _getOptimisticOracle();
        requestTime0 = block.timestamp;
        require(
        
            oracle.hasPrice(address(this), _priceIdentifier, requestTime0, ""),
            "Unresolved oracle price"
        );
        int256 oraclePrice = oracle.settleAndGetPrice(_priceIdentifier, requestTime0, "");

        // For simplicity we don't want to deal with negative prices.
        if (oraclePrice < 0) {
            oraclePrice = 0;
        }
        return uint256(oraclePrice);
    }

    function requestData(bytes32 _identifier) external returns(uint256) {
        requestTime = block.timestamp;
        IERC20 bondCurrency = IERC20(0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14); //weth sepolia
        uint256 reward = 0;


        umaOraclePrice =  oo.requestPrice(_identifier, requestTime, "", bondCurrency, reward);
       
        oo.setCustomLiveness(_identifier, requestTime, "", 30);

        return umaOraclePrice;



    }


}