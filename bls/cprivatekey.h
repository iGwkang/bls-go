#pragma once
#include <stdint.h>
#include "celements.h"

const static int PRIVATE_KEY_SIZE = 32;

typedef void * CPrivateKey;

CPrivateKey PrivateKeyFromBytes(uint8_t bytes[PRIVATE_KEY_SIZE]);
void FreePrivateKey(CPrivateKey);

// return data size == PRIVATE_KEY_SIZE
uint8_t *PrivateKeySerialize(CPrivateKey);

CPrivateKey PrivateKeyCopy(CPrivateKey);

// 返回值不必释放
CG1Element PrivateKeyGetG1(CPrivateKey);
CG2Element PrivateKeyGetG2(CPrivateKey);

CPrivateKey PrivateKeyAggregate(CPrivateKey *, int count);

bool EqualPrivateKey(CPrivateKey, CPrivateKey);

CG1Element MulG1PrivateKey(CPrivateKey, CG1Element);

CG2Element MulG2PrivateKey(CPrivateKey, CG2Element);

