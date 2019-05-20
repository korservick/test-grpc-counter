# Final stage
FROM alpine
EXPOSE 50051
WORKDIR /
COPY --from=korservick/grpc_counter:compile /counter /
ENTRYPOINT ["/counter"]