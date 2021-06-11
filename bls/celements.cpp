extern "C" {
#include "celements.h"
}
#include "elements.hpp"

CG1Element NewG1Element()
{
	return new bls::G1Element();
}

void FreeG1Element(CG1Element p)
{
	delete (bls::G1Element*)p;
}

CG1Element G1ElementCopy(CG1Element p)
{
	return new bls::G1Element(*(bls::G1Element*)p);
}

CG1Element G1ElementFromBytes(const uint8_t *data, int len)
{
	bls::Bytes bytes(data, len);
	return new bls::G1Element(bls::G1Element::FromBytes(bytes));
}

CG1Element G1ElementFromMessage(const uint8_t *data, int len, const uint8_t *dst_data, int dst_len)
{
	bls::Bytes bytes(data, len);
	return new bls::G1Element(bls::G1Element::FromMessage(bytes, dst_data, dst_len));
}

CG1Element G1ElementGenerator()
{
	return new bls::G1Element(bls::G1Element::Generator());
}

CG1Element G1ElementNegate(CG1Element p)
{
	return new bls::G1Element(((bls::G1Element*)p)->Negate());
}

uint32_t G1ElementGetFingerprint(CG1Element p)
{
	return ((bls::G1Element*)p)->GetFingerprint();
}

void G1ElementSerialize(CG1Element p, uint8_t **out_data, uint32_t *out_size)
{
	std::vector<uint8_t> data = ((bls::G1Element*)p)->Serialize();
	*out_size = data.size();
	*out_data = (uint8_t *)malloc(*out_size);
	memcpy(*out_data, &data[0], data.size());
}

bool EqualG1Element(CG1Element lhs, CG1Element rhs)
{
	return *(bls::G1Element*)lhs == *(bls::G1Element*)rhs;
}

CG1Element AddG1Element(CG1Element self, CG1Element other)
{
	bls::G1Element newEle = *(bls::G1Element*)self + *(bls::G1Element*)other;

	return new bls::G1Element(newEle);
}

CG1Element MulG1Element(CG1Element self, const bn_t * k)
{
	bls::G1Element newEle = *(bls::G1Element*)self * (*k);
	return new bls::G1Element(newEle);
}

// ------------------------------------------------------------------------------- //
CG2Element NewG2Element()
{
	return new bls::G2Element();
}

void FreeG2Element(CG2Element p)
{
	delete (bls::G2Element*)p;
}

CG2Element G2ElementCopy(CG2Element p)
{
	return new bls::G2Element(*(bls::G2Element*)p);
}

CG2Element G2ElementFromBytes(const uint8_t *data, int len)
{
	bls::Bytes bytes(data, len);
	return new bls::G2Element(bls::G2Element::FromBytes(bytes));
}

CG2Element G2ElementFromMessage(const uint8_t *data, int len, const uint8_t *dst_data, int dst_len)
{
	bls::Bytes bytes(data, len);
	return new bls::G2Element(bls::G2Element::FromMessage(bytes, dst_data, dst_len));
}

CG2Element G2ElementGenerator()
{
	return new bls::G2Element(bls::G2Element::Generator());
}

CG2Element G2ElementNegate(CG2Element p)
{
	return new bls::G2Element(((bls::G2Element*)p)->Negate());
}

void G2ElementSerialize(CG2Element p, uint8_t **out_data, uint32_t *out_size)
{
	std::vector<uint8_t> data = ((bls::G2Element*)p)->Serialize();
	*out_size = data.size();
	*out_data = (uint8_t *)malloc(*out_size);
	memcpy(*out_data, &data[0], data.size());
}

bool EqualG2Element(CG2Element lhs, CG2Element rhs)
{
	return *(bls::G2Element*)lhs == *(bls::G2Element*)rhs;
}

CG2Element AddG2Element(CG2Element self, CG2Element other)
{
	bls::G2Element newEle = *(bls::G2Element*)self + *(bls::G2Element*)other;

	return new bls::G2Element(newEle);
}

CG2Element MulG2Element(CG2Element self, const bn_t * k)
{
	bls::G2Element newEle = *(bls::G2Element*)self * (*k);
	return new bls::G2Element(newEle);
}

