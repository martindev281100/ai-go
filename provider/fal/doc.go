// Package fal implements fal.ai video generation over the fal queue API
// (https://queue.fal.run). Any endpoint id is accepted; model input is built
// from the request plus [VideoOptions.RawInput].
//
// Besides the blocking [VideoModel.GenerateVideo], a [VideoModel] exposes a
// durable split for applications that persist requests: [VideoModel.SubmitReply],
// [VideoModel.StatusReply], and [VideoModel.ResultReply] perform exactly one
// untranslated HTTP exchange each, and the pure decoders [DecodeSubmission],
// [DecodeStatus], and [DecodeVideoResult] classify those replies as accepted,
// rejected, ambiguous, pending, completed, transient, or contract failures.
package fal
