/*
Garrett R. Fairburn
10/21/16
Simple Block Cipher I.
*/

#include <stdint.h>
#include <stdlib.h>

// Base invert-shift-xor engine. (Cryptographic primitive for SBCI.)
void ISXEngine(const uint8_t * mask, const uint8_t * previous, uint8_t * current) {

	// Set state.
	((uint64_t*)current)[0] = ((uint64_t*)previous)[0];

	// Set modifier.
	uint8_t modifier[8];
	((uint64_t*)modifier)[0] = ((uint64_t*)previous)[0] ^ ((uint64_t*)mask)[0];

	// Store shift information.
	uint8_t shift = 0;

	// Invert.
	for (uint8_t i = 0; i < 8; i++) {

		// Get locations.
		uint8_t locations[2] = { (uint8_t)(modifier[i] & 0x0F), (uint8_t)((modifier[i] & 0xF0) >> 4) };

		// Perform inversions.
		for (uint8_t j = 0; j < 2; j++) {

			if (locations[j] < 8) {

				// Column
				for (uint8_t k = 0; k < 8; k++) {

					current[k] ^= 0b10000000 >> locations[j];
				}
			}
			else {

				// Row.
				current[locations[j] - 8] ^= 0xFF;
			}
		}
	}

	// Compress state to get shift info.
	for (uint8_t i = 0; i < 8; i++) {

		shift ^= current[i];
	}
	shift ^= shift >> 4;

	// Shift state.
	for (uint8_t i = 0; i <= (shift & 0b00000111); i++) {

		// Overflow byte.
		uint8_t overflow = 0;

		// Choose direction.
		if ((shift & 0b00001000) == 0) {

			// Shift right.
			overflow = current[0] & 0b00000001;
			((uint64_t*)current)[0] >>= 1;
			current[7] |= overflow << 7;
		}
		else {

			// Shift left.
			overflow = current[7] & 0b10000000;
			((uint64_t*)current)[0] <<= 1;
			current[0] |= overflow >> 7;
		}
	}

	// Xor state with mask.
	((uint64_t*)current)[0] ^= ((uint64_t*)mask)[0];

	return;
}

// Use ISXEngine to generate a byte sequence.
uint8_t* ISXSequence(const uint8_t* seed, uint32_t seedSize, uint32_t arraySize) {

	// Resize array for sequence.
	uint8_t * workingArray = malloc(arraySize);

	// Temporary sequence holder.
	uint32_t tempArraySize = arraySize;
	if (tempArraySize % 8 != 0) {

		tempArraySize += 8 - (tempArraySize % 8);
	}
	uint8_t * tempArray = malloc(tempArraySize);

	// Temporary seed holder.
	uint32_t tempSeedSize = seedSize;
	if (tempSeedSize % 8 != 0) {

		tempSeedSize += 8 - (tempSeedSize % 8);
	}
	uint8_t * tempSeed = malloc(tempSeedSize);

	// Copy seed to tempSeed.
	for (uint64_t i = 0; i < seedSize; i++) {

		tempSeed[i] = seed[i];
	}

	// Value for seeding engine.
	uint64_t IniV;

	// Generate sequences.
	for (uint64_t i = 0; i < seedSize; i += 8) {

		// Seed engine.
		ISXEngine(tempSeed + i, tempSeed + i, (uint8_t*)&IniV);

		// Generate first value.
		ISXEngine(tempSeed + i, (uint8_t*)&IniV, tempArray);

		// Generate subsequent values.
		for (uint64_t j = 8; j < tempArraySize; j += 8) {

			ISXEngine(tempSeed + i, tempArray + j - 8, tempArray + j);
		}

		// Xor tempArray with workingArray.
		for (uint64_t j = 0; j < arraySize; j++) {

			workingArray[j] ^= tempArray[j];
		}
	}

	return workingArray;
}
