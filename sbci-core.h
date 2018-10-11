/*
Garrett R. Fairburn
10/21/16
Simple Block Cipher I.
*/

#pragma once

#include <stdint.h>

// Base invert-shift-xor engine. (Cryptographic primitive for SBCI.)
void ISXEngine(const uint8_t*, const uint8_t*, uint8_t*);

// Use ISXEngine to generate a quadword sequence.
uint8_t* ISXSequence(const uint8_t*, uint32_t, uint32_t);
