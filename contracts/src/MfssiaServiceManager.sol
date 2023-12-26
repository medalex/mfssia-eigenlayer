// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@eigenlayer/contracts/libraries/BytesLib.sol";
import "./IMfssiaTaskManager.sol";
import "@eigenlayer-middleware/src/ServiceManagerBase.sol";

/**
 * @title Primary entrypoint for procuring services from Mfssia.
 * @author Layr Labs, Inc.
 */
contract MfssiaServiceManager is ServiceManagerBase {
    using BytesLib for bytes;

    IMfssiaTaskManager
        public immutable mfssiaTaskManager;

    /// @notice when applied to a function, ensures that the function is only callable by the `registryCoordinator`.
    modifier onlyMfssiaTaskManager() {
        require(
            msg.sender == address(mfssiaTaskManager),
            "onlyMfssiaTaskManager: not from mfssia task manager"
        );
        _;
    }

    constructor(
        IBLSRegistryCoordinatorWithIndices _registryCoordinator,
        ISlasher _slasher,
        IMfssiaTaskManager _mfssiaTaskManager
    ) ServiceManagerBase(_registryCoordinator, _slasher) {
        mfssiaTaskManager = _mfssiaTaskManager;
    }

    /// @notice Called in the event of challenge resolution, in order to forward a call to the Slasher, which 'freezes' the `operator`.
    /// @dev The Slasher contract is under active development and its interface expected to change.
    ///      We recommend writing slashing logic without integrating with the Slasher at this point in time.
    function freezeOperator(
        address operatorAddr
    ) external override onlyMfssiaTaskManager {
        // slasher.freezeOperator(operatorAddr);
    }
}
