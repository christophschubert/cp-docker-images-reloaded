FROM registry.access.redhat.com/ubi8/go-toolset as builder

WORKDIR /build
COPY ub/ub.go .
RUN go build -o ub .

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4-205
COPY --from=builder /build/ub /usr/bin
