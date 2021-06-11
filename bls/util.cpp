extern "C" {
#include "util.h"
#include "relic.h"
#include "relic_md.h"
}
#include "util.hpp"

void Hash256(uint8_t* output, const uint8_t* message, int messageLen)
{
	bls::Util::Hash256(output, message, messageLen);
}

