// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/MfssiaServiceManager.sol" as incsqsm;
import {MfssiaTaskManager} from "../src/MfssiaTaskManager.sol";
import {BLSMockAVSDeployer} from "@eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract MfssiaTaskManagerTest is BLSMockAVSDeployer {
    incsqsm.MfssiaServiceManager sm;
    incsqsm.MfssiaServiceManager smImplementation;
    MfssiaTaskManager tm;
    MfssiaTaskManager tmImplementation;

    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    address aggregator =
        address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    address generator =
        address(uint160(uint256(keccak256(abi.encodePacked("generator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();

        tmImplementation = new MfssiaTaskManager(
            incsqsm.IBLSRegistryCoordinatorWithIndices(
                address(registryCoordinator)
            ),
            TASK_RESPONSE_WINDOW_BLOCK
        );

        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        tm = MfssiaTaskManager(
            address(
                new TransparentUpgradeableProxy(
                    address(tmImplementation),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        tm.initialize.selector,
                        pauserRegistry,
                        serviceManagerOwner,
                        aggregator,
                        generator
                    )
                )
            )
        );
    }

    function testCreateNewTask() public {
        bytes memory quorumNumbers = new bytes(0);
        cheats.prank(generator, generator);
        tm.createNewTask("3636a05186cdf92f726fc5c9c2c9f961ae9b9c48f7ab841723fc539c10fe7f77","3636a05186cdf92f726fc5c9c2c9f961ae9b9c48f7ab841723fc539c10fe7f75","3636a05186cdf92f726fc5c9c2c9f961ae9b9c48f7ab841723fc539c10fe7f77", 100, quorumNumbers);
        assertEq(tm.latestTaskNum(), 1);
    }
}
